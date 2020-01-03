package services

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
func (c *HuobiRestClient) GetAccounts() models.AccountsReturn {
	accountsReturn := models.AccountsReturn{}

	strRequest := "/v1/account/accounts"
	jsonAccountsReturn := utils.ApiKeyGet(c.Config, make(map[string]string), strRequest)
	_ = json.Unmarshal([]byte(jsonAccountsReturn), &accountsReturn)

	return accountsReturn
}

// 根据账户ID查询账户余额
// nAccountID: 账户ID, 不知道的话可以通过GetAccounts()获取, 可以只现货账户, C2C账户, 期货账户
// return: BalanceReturn对象
func (c *HuobiRestClient) GetAccountBalance(strAccountID string) (models.BalanceReturn, error) {
	balanceReturn := models.BalanceReturn{}

	strRequest := fmt.Sprintf("/v1/account/accounts/%s/balance", strAccountID)
	jsonBalanceReturn := utils.ApiKeyGet(c.Config, make(map[string]string), strRequest)
	fmt.Println(jsonBalanceReturn)
	err := json.Unmarshal([]byte(jsonBalanceReturn), &balanceReturn)
	if err != nil {
		return balanceReturn, err
	}
	return balanceReturn, nil
}

// 根据账户ID查询账户余额
// nAccountID: 账户ID, 不知道的话可以通过GetAccounts()获取, 可以只现货账户, C2C账户, 期货账户
// return: BalanceReturn对象
func (c *HuobiRestClient) GetAccountBalance(strAccountID string) (models.BalanceReturn, error) {
	balanceReturn := models.BalanceReturn{}

	strRequest := fmt.Sprintf("/v1/account/accounts/%s/balance", strAccountID)
	jsonBalanceReturn := utils.ApiKeyGet(c.Config, make(map[string]string), strRequest)
	fmt.Println(jsonBalanceReturn)
	err := json.Unmarshal([]byte(jsonBalanceReturn), &balanceReturn)
	if err != nil {
		return balanceReturn, err
	}
	return balanceReturn, nil
}

// ------------------------------------------------------------------------------------------
// 资金API

// 账号资金划转
// TransferRequestParams: 下单信息
// return: TransferReturn 对象
func (c *HuobiRestClient) Transfer(transferRequestParams models.TransferRequestParams) models.TransferReturn {
	transferReturn := models.TransferReturn{}

	mapParams := make(map[string]string)
	mapParams["sub-uid"] = transferRequestParams.SubUID
	mapParams["currency"] = transferRequestParams.Currency
	mapParams["amount"] = transferRequestParams.Amount
	mapParams["type"] = string(transferRequestParams.Type)

	strRequest := "/v1/subuser/transfer"
	jsonTransferReturn := utils.ApiKeyPost(c.Config, mapParams, strRequest)
	_ = json.Unmarshal([]byte(jsonTransferReturn), &transferReturn)
	// master-transfer-in（子账号划转给母账户虚拟币）
	// master-transfer-out （母账户划转给子账号虚拟币）
	// master-point-transfer-in （子账号划转给母账户点卡）
	// master-point-transfer-out（母账户划转给子账号点卡）

	return transferReturn
}

// 虚拟币提现
// WithdrawRequestParams: 提现信息
// return: WithdrawReturn 对象
func (c *HuobiRestClient) Withdraw(params models.WithdrawRequestParams) models.WithdrawReturn {
	withdrawReturn := models.WithdrawReturn{}

	mapParams := make(map[string]string)
	mapParams["address"] = params.Address
	mapParams["currency"] = params.Currency
	mapParams["amount"] = params.Amount

	strRequest := "/v1/dw/withdraw/api/create"
	jsonWithdrawReturn := utils.ApiKeyPost(c.Config, mapParams, strRequest)
	_ = json.Unmarshal([]byte(jsonWithdrawReturn), &withdrawReturn)

	return withdrawReturn
}
