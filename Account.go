package huobi

import (
	"encoding/json"
	"fmt"

	"github.com/shaodan/go-huobi/models"
	"github.com/shaodan/go-huobi/utils"
)

// ------------------------------------------------------------------------------------------
// 用户资产API

// 查询当前用户的所有账户, 根据包含的私钥查询
// return: AccountsReturn对象
func (c *RestClient) GetAccounts() models.AccountsReturn {
	accountsReturn := models.AccountsReturn{}

	strRequest := "/v1/account/accounts"
	jsonAccountsReturn := utils.ApiKeyGet(c.Config, make(map[string]string), strRequest)
	_ = json.Unmarshal([]byte(jsonAccountsReturn), &accountsReturn)

	return accountsReturn
}

// 根据账户ID查询账户余额
// nAccountID: 账户ID, 不知道的话可以通过GetAccounts()获取, 可以只现货账户, C2C账户, 期货账户
// return: BalanceReturn对象
func (c *RestClient) GetAccountBalance(strAccountID string) models.BalanceReturn {
	balanceReturn := models.BalanceReturn{}

	strRequest := fmt.Sprintf("/v1/account/accounts/%s/balance", strAccountID)
	jsonBalanceReturn := utils.ApiKeyGet(c.Config, make(map[string]string), strRequest)
	_ = json.Unmarshal([]byte(jsonBalanceReturn), &balanceReturn)

	return balanceReturn
}
