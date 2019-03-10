package services

import (
	"log"
	"testing"
)

func TestTransfer(t *testing.T) {
	testLogin()

	accounts := testClient.GetAccounts()
	log.Printf("%+v", accounts)
}
