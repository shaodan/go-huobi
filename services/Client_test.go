package services

var testClient = NewHuobiRestClient("api.huobi.pro")

func testLogin() {
	accessKey := "your access key here"
	secretKey := "your secret key here"
	testClient.Login(accessKey, secretKey)
}
