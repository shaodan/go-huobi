package huobi

import (
	"testing"
)

func TestRestClient_GetAccounts(t *testing.T) {
	testLogin()
	account := testClient.GetAccounts()
	t.Log("Account: ", account)
}