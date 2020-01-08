package huobi

import (
	"log"
	"testing"

	"github.com/shaodan/go-huobi/models"
)

func TestGetLoanOrders(t *testing.T) {
	testLogin()

	accounts := testClient.GetLoanOrders(models.GetLoanRequestParams{Symbol: "bchbtc"})

	log.Printf("%+v", accounts)
}
