package huobi

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/shaodan/go-huobi/models"
)

func TestPlaceOrder(t *testing.T) {
	testLogin()

	account := testClient.GetAccounts()

	fmt.Println("Account: ", account)

	if strings.Compare(account.Status, "ok") == 0 {
		accounts := account.Data

		if len(accounts) >= 1 {
			var account models.SubAccount

			for _, entry := range accounts {
				if entry.Type == "spot" {
					account = entry
					break
				}
			}

			if 0 != account.ID {
				var placeParams models.PlaceRequestParams
				placeParams.AccountID = strconv.Itoa(int(account.ID))
				placeParams.Amount = "1.0"
				placeParams.Price = "5721"
				placeParams.Source = "api"
				placeParams.Symbol = "btcusdt"
				placeParams.Type = "sell-limit"

				fmt.Println("Place order with: ", placeParams)
				placeReturn := testClient.Place(placeParams)
				if placeReturn.Status == "ok" {
					fmt.Println("Place return: ", placeReturn.Data)
				} else {
					t.Errorf("Place error: %s", placeReturn.ErrMsg)
				}
			} else {
				t.Error("account is nil")
			}

		}

	} else {
		t.Error(account.ErrMsg)
	}
}
