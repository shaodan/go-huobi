package config

type HuobiRestConfig struct {
	Logon     bool
	AccessKey string
	SecretKey string

	// default to be disabled, please DON'T enable it unless it's officially announced.
	EnablePrivateSignature bool

	// generated the key by: openssl ecparam -name prime256v1 -genkey -noout -out privatekey.pem
	// only required when Private Signature is enabled
	// todo: replace with your own PrivateKey from privatekey.pem
	PrivateKeyPrime256 string

	// API请求地址, 不要带最后的/
	//todo: replace with real URLs and HostName
	MarketUrl string
	TradeUrl  string
	HostName  string
}
