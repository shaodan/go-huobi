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
