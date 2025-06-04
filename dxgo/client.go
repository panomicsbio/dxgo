package dxgo

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/avast/retry-go/v4"
)

type DXSecurityContext struct {
	AuthToken     string `json:"auth_token"`
	AuthTokenType string `json:"auth_token_type"`
}

type DXClientConfig struct {
	PublicApiOrigin string

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
		PublicApiOrigin:   "https://api.dnanexus.com",
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

func (c *DXClient) getBaseOrigin() string {
	return fmt.Sprintf("%s://%s:%s", c.config.ApiServerProtocol, c.config.ApiServerHost, c.config.ApiServerPort)
}

func (c *DXClient) DoInto(ctx context.Context, uri string, input any, output any) error {
	data, err := c.retryableRequest(ctx, uri, input, nil)
	if err != nil {
		return fmt.Errorf("making retryable request: %w", err)
	}

	if output != nil {
		err = json.Unmarshal(data, output)
		if err != nil {
			return fmt.Errorf("unmarshalling data: %w", err)
		}
	}

	return nil
}

type DXClientOptions struct {
	PublicApi bool
	Headers   map[string]string
}

func (c *DXClient) DoIntoWithOptions(ctx context.Context, uri string, input any, output any, options *DXClientOptions) error {
	data, err := c.retryableRequest(ctx, uri, input, options)
	if err != nil {
		return fmt.Errorf("making retryable request: %w", err)
	}

	if output != nil {
		err = json.Unmarshal(data, output)
		if err != nil {
			return fmt.Errorf("unmarshalling data: %w", err)
		}
	}

	return nil
}

func (c *DXClient) retryableRequest(ctx context.Context, uri string, input any, options *DXClientOptions) ([]byte, error) {
	var resp []byte
	err := retry.Do(func() error {
		var err error
		resp, err = c.request(ctx, uri, input, options)
		if err != nil {
			slog.Log(ctx, slog.LevelError, "error making request", slog.Any("err", err))
		}
		return err
	}, retry.DelayType(retryDelay), retry.Attempts(c.config.MaxRetries))
	if err != nil {
		return nil, fmt.Errorf("doing retry on dxclient request: %w", err)
	}

	return resp, nil
}

func (c *DXClient) request(ctx context.Context, uri string, input any, options *DXClientOptions) ([]byte, error) {
	postUrl := fmt.Sprintf("%s%s", c.getBaseOrigin(), uri)
	if options.PublicApi {
		postUrl = fmt.Sprintf("%s%s", c.config.PublicApiOrigin, uri)
	}

	data, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("marshalling request input: %w", err)
	}

	r, err := http.NewRequestWithContext(ctx, "POST", postUrl, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("creating http request: %w", err)
	}

	r.Header.Add("Authorization", fmt.Sprintf("%s %s", c.config.DXSecurityContext.AuthTokenType, c.config.DXSecurityContext.AuthToken))
	r.Header.Add("Content-Type", "application/json")
	for k, v := range options.Headers {
		r.Header.Add(k, v)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("making http request: %w", err)
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			slog.LogAttrs(ctx, slog.LevelError, "closing response body", slog.Any("err", err))
		}
	}()

	if res.StatusCode == http.StatusServiceUnavailable {
		return nil, RetryAfterError{response: *res}
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
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

func (c *DXClient) GetApiServerProtocol() string {
	return c.config.ApiServerProtocol
}

func (c *DXClient) GetApiServerHost() string {
	return c.config.ApiServerHost
}

func (c *DXClient) GetApiServerPort() string {
	return c.config.ApiServerPort
}
