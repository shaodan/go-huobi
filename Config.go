package huobi

type RestConfig struct {
	Logon     bool
	AccessKey string
	SecretKey string

	// default to be disabled, please DON'T enable it unless it's officially announced.
	EnablePrivateSignature bool

	// generated the key by: openssl ecparam -name prime256v1 -genkey -noout -out privatekey.pem
	// only required when Private Signature is enabled
	// replace with your own PrivateKey from privatekey.pem
	PrivateKeyPrime256 string

	// API请求地址, 不要带最后的/
	MarketUrl string
	TradeUrl  string
	HostName  string
}
