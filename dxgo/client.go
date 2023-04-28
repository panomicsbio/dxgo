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
	"os"
	"time"
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

func (c *DXClient) retryableRequest(uri string, input interface{}) ([]byte, error) {
	var resp []byte
	err := retry.Do(func() error {
		postUrl := fmt.Sprintf("%s%s", c.getBaseEndpoint(), uri)
		data, err := json.Marshal(input)
		if err != nil {
			return err
		}
		r, err := http.NewRequest("POST", postUrl, bytes.NewReader(data))
		if err != nil {
			return err
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
		if res.StatusCode == 503 {
			return RetryAfterError{response: *res}
		}
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
	}), retry.Attempts(c.config.MaxRetries))
	return resp, err
}
