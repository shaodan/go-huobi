package services

import (
	"testing"
)

func Test_getSymbols(t *testing.T) {
	symbols := testClient.GetSymbols()
	if symbols.Status != "ok" {
		t.Error("failed to get symbols")
	} else {
		t.Logf("toal symbols: %v", len(symbols.Data))
	}
}

func Test_GetAccountBalance(t *testing.T) {
	testLogin()
	testClient.GetAccountBalance("6003956")
}

func TestHuobiRestClient_GetKLine(t *testing.T) {
	testLogin()
	testClient.GetKLine("bchbtc", "1min", 2000)
}
