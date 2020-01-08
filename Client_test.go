package huobi

var testClient = NewRestClient("api.huobi.pro")

func testLogin() {
	accessKey := "your access key here"
	secretKey := "your secret key here"
	testClient.Login(accessKey, secretKey)
}
