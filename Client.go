package huobi

import (
	"fmt"
)

type RestClient struct {
	Config *RestConfig
}

func NewRestClient(hostname string) *RestClient {
	url := fmt.Sprintf("https://%s", hostname)
	restConfig := &RestConfig{
		MarketUrl: url,
		TradeUrl:  url,
		HostName:  hostname,
	}
	return &RestClient{
		Config: restConfig,
	}
}

func (c *RestClient) Login(accessKey, secretKey string) *RestClient {
	cfg := c.Config
	cfg.AccessKey = accessKey
	cfg.SecretKey = secretKey
	cfg.EnablePrivateSignature = false
	cfg.Logon = true
	return c
}
