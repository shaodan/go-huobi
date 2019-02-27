package services

import (
	"fmt"
	"github.com/shaodan/huobi-rest/config"
)

type HuobiRestClient struct {
	Config *config.HuobiRestConfig
}

func NewHuobiRestClient(hostname string) *HuobiRestClient {
	url := fmt.Sprintf("https://%s", hostname)
	config := &config.HuobiRestConfig{
		MarketUrl: url,
		TradeUrl:  url,
		HostName:  hostname,
	}
	return &HuobiRestClient{
		Config: config,
	}
}

func (c *HuobiRestClient) Login(accessKey, secretKey string) {
	cfg := c.Config
	cfg.AccessKey = accessKey
	cfg.SecretKey = secretKey
	cfg.EnablePrivateSignature = false
	cfg.Logon = true
}
