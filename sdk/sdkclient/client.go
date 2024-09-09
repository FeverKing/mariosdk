package sdkclient

import (
	"mariosdk/sdk/sdkreq"
)

type Client interface {
	setConfig(config Config)
	Auth() error
	NewClient() Client
}

type DefaultClient struct {
	Config    Config
	apiClient *sdkreq.ApiClient
}

func (c *DefaultClient) setConfig(config Config) {
	c.Config = config
}

func NewClient() *DefaultClient {
	dc := &DefaultClient{}
	return dc
}
