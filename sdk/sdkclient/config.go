package sdkclient

type Config struct {
	AccessKey string   `json:"accessKey"`
	SecretKey string   `json:"secretKey"`
	Endpoints []string `json:"endpoints"`
}

func (c *Config) SetAccessKey(accessKey string) {
	c.AccessKey = accessKey
}

func (c *Config) SetSecretKey(secretKey string) {
	c.SecretKey = secretKey
}

func (c *Config) AddEndpoint(endpoint string) {
	c.Endpoints = append(c.Endpoints, endpoint)
}
