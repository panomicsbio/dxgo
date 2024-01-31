package dxgo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/avast/retry-go"
)

type DXSecurityContext struct {
	AuthToken     string `json:"auth_token"`
	AuthTokenType string `json:"auth_token_type"`
}

type DXClientConfig struct {
	ApiServerProtocol string
	ApiServerHost     string
	ApiServerPort     string
	DXSecurityContext *DXSecurityContext
	MaxRetries        uint
}

type DXClient struct {
	config *DXClientConfig
}

func NewClient(maxRetries uint) (*DXClient, error) {
	sc := new(DXSecurityContext)
	err := json.Unmarshal([]byte(os.Getenv("DX_SECURITY_CONTEXT")), &sc)
	if err != nil {
		return nil, err
	}
	return &DXClient{config: &DXClientConfig{
		ApiServerProtocol: os.Getenv("DX_APISERVER_PROTOCOL"),
		ApiServerHost:     os.Getenv("DX_APISERVER_HOST"),
		ApiServerPort:     os.Getenv("DX_APISERVER_PORT"),
		DXSecurityContext: sc,
		MaxRetries:        maxRetries,
	}}, nil
}

func NewClientWithConfig(config *DXClientConfig) *DXClient {
	return &DXClient{config: config}
}

func (c *DXClient) getBaseEndpoint() string {
	return fmt.Sprintf("%s://%s:%s", c.config.ApiServerProtocol, c.config.ApiServerHost, c.config.ApiServerPort)
}

func (c *DXClient) DoInto(uri string, input any, output any) error {
	data, err := c.retryableRequest(uri, input)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, output)
}

func (c *DXClient) retryableRequest(uri string, input interface{}) ([]byte, error) {
	var resp []byte
	err := retry.Do(func() error {
		var err error
		resp, err = c.request(uri, input)
		return err
	}, retry.DelayType(retryDelay), retry.Attempts(c.config.MaxRetries))

	return resp, err
}

func (c *DXClient) request(uri string, input interface{}) ([]byte, error) {
	postUrl := fmt.Sprintf("%s%s", c.getBaseEndpoint(), uri)
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", postUrl, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", c.config.DXSecurityContext.AuthTokenType, c.config.DXSecurityContext.AuthToken))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	if res.StatusCode == 503 {
		return nil, RetryAfterError{response: *res}
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func retryDelay(n uint, err error, config *retry.Config) time.Duration {
	var e RetryAfterError
	if errors.As(err, &e) {
		if t, err := ParseRetryAfter(e.response.Header.Get("Retry-After")); err == nil {
			return time.Until(t)
		}
	}

	return retry.BackOffDelay(n, err, config)
}

func (c *DXClient) GetMaxRetries() uint {
	return c.config.MaxRetries
}
