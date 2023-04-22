package dxgo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/aereal/go-httpretryafter"
	"github.com/avast/retry-go"
	"io"
	"net/http"
	"time"
)

type RetryAfterError struct {
	response http.Response
}

func (err RetryAfterError) Error() string {
	return fmt.Sprintf(
		"Request to %s fail %s (%d)",
		err.response.Request.RequestURI,
		err.response.Status,
		err.response.StatusCode,
	)
}

func (c *DXClient) request(path string, input interface{}) ([]byte, error) {
	var resp []byte
	err := retry.Do(func() error {
		postUrl := fmt.Sprintf("%s%s", c.getBaseEndpoint(), path)
		data, err := json.Marshal(input)
		if err != nil {
			return err
		}
		r, err := http.NewRequest("POST", postUrl, bytes.NewReader(data))
		if err != nil {
			panic(err)
		}
		r.Header.Add("Authorization", fmt.Sprintf("%s %s", c.config.DXSecurityContext.AuthTokenType, c.config.DXSecurityContext.AuthToken))
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		res, err := client.Do(r)
		if err != nil {
			return err
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		resp, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return nil
	}, retry.DelayType(func(n uint, err error, config *retry.Config) time.Duration {
		switch e := err.(type) {
		case RetryAfterError:
			if t, err := httpretryafter.Parse(e.response.Header.Get("Retry-After")); err == nil {
				return time.Until(t)
			}
		}
		return retry.BackOffDelay(n, err, config)
	}))
	return resp, err
}
