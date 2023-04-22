package dxgo

import (
	"encoding/json"
	"fmt"
	"os"
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
}

type DXClient struct {
	config *DXClientConfig
}

func NewClient() (*DXClient, error) {
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
	}}, nil
}

func NewClientWithConfig(config *DXClientConfig) *DXClient {
	return &DXClient{config: config}
}

func (c *DXClient) getBaseEndpoint() string {
	return fmt.Sprintf("%://%s:%s", c.config.ApiServerProtocol, c.config.ApiServerHost, c.config.ApiServerPort)
}
