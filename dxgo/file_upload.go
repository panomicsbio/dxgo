package dxgo

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"iter"
	"net/http"
	"time"

	retryRequest "github.com/avast/retry-go/v4"
	"github.com/panomicsbio/utils"
)

type Part interface {
	Reader() io.Reader
	Size() uint64
}

// Determines the recommended size for the part upload. If desiredPartSize is negative, the value will be determined automatically.
// The minimum part size is usually 5mb, where the maximum part size is hundreds of gb
func (c *DXClient) DeterminePartSize(ctx context.Context, projectId string, desiredPartSize int64) (int64, error) {
	pi, err := c.ProjectDescribe(ctx, projectId, ProjectDescribeInput{
		Fields: map[string]bool{
			"fileUploadParameters": true,
		},
	})
	if err != nil {
		return -1, fmt.Errorf("describing project: %w", err)
	}

	if pi.Error != nil {
		return -1, fmt.Errorf("describing project output error: %w", pi.Error)
	}

	minPartSize := int64(pi.FileUploadParameters["minimumPartSize"].(float64))
	maxPartSize := int64(pi.FileUploadParameters["maximumPartSize"].(float64))
	if desiredPartSize < 0 {
		return minPartSize, nil
	}

	if desiredPartSize < minPartSize {
		return -1, fmt.Errorf("Cannot upload parts of size %d, the minimum part size is %d", desiredPartSize, minPartSize)
	}

	return min(desiredPartSize, maxPartSize), nil
}

func (c *DXClient) DoMultipartUpload(ctx context.Context, projectId, folder, objectName string, parts iter.Seq2[Part, error], waitForClosing bool) error {
	fn, err := c.FileNew(ctx, FileNewInput{
		Project: projectId,
		Folder:  utils.SetPrefixSlash(folder),
		Parent:  true,
		Name:    objectName,
	})
	if err != nil {
		return fmt.Errorf("file new-ing: %w", err)
	}

	if fn.Error != nil {
		return fmt.Errorf("file new-ing output error: %w", fn.Error)
	}

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	partNumber := 1
	for part, err := range parts {
		if err != nil {
			return fmt.Errorf("could not get chunk: %w", err)
		}

		buf := make([]byte, part.Size())
		n, err := io.ReadFull(part.Reader(), buf)
		if err != nil {
			return fmt.Errorf("failed to read: %w", err)
		}

		if n != int(part.Size()) {
			return fmt.Errorf("expected to read %d, but only read %d", part.Size(), n)
		}

		md5 := utils.CreateMd5Hash(buf)
		fu, err := c.FileUpload(ctx, FileUploadInput{
			ID:    fn.ID,
			Size:  n,
			MD5:   md5,
			Index: partNumber,
		})
		if err != nil {
			return fmt.Errorf("uploading file: %w", err)
		}

		if fu.Error != nil {
			return fmt.Errorf("uploading file output error: %w", fu.Error)
		}

		body := bytes.NewBuffer(buf)
		err = retryRequest.Do(
			func() error {
				req, e := http.NewRequestWithContext(ctx, http.MethodPut, fu.URL, body)
				if e != nil {
					return fmt.Errorf("creating http request: %w", err)
				}

				for k, v := range fu.Headers {
					req.Header.Add(k, v)
				}

				resp, e := client.Do(req)
				if e != nil {
					return fmt.Errorf("doing http request: %w", err)
				}

				if resp.StatusCode < 200 || resp.StatusCode > 299 {
					return fmt.Errorf("non-2xx response: %d - %s", resp.StatusCode, resp.Status)
				}

				return nil
			},
			retryRequest.DelayType(func(n uint, err error, config *retryRequest.Config) time.Duration {
				return retryRequest.BackOffDelay(n, err, config)
			}),
			retryRequest.Attempts(c.GetMaxRetries()),
		)
		if err != nil {
			return fmt.Errorf("retrying request: %w", err)
		}

		partNumber += 1
	}

	fc, err := c.FileClose(ctx, FileCloseInput{ID: fn.ID})
	if err != nil {
		return fmt.Errorf("closing file: %w", err)
	}
	if fc.Error != nil {
		return fmt.Errorf("closing file output error: %w", fc.Error)
	}

	if !waitForClosing {
		return nil
	}

	for {
		fd, err := c.FileDescribe(ctx, FileDescribeInput{ID: fn.ID})
		if err != nil {
			return fmt.Errorf("describing file to determine if it's uploaded: %w", err)
		}

		if fd.Error != nil {
			return fmt.Errorf("describing file to determine if it's uploaded output error: %w", fd.Error)
		}

		if fd.State == "closed" {
			return nil
		}

		time.Sleep(2 * time.Second)
	}
}
