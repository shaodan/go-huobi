package services

import (
	"encoding/json"
	"fmt"

	"github.com/shaodan/go-huobi/models"
	"github.com/shaodan/go-huobi/utils"
)

//------------------------------------------------------------------------------------------
// 借贷API

// 查询借贷订单
// GetLoanRequestParams: 查询信息
// return: GetLoanReturn对象
func (c *RestClient) GetLoanOrders(getLoanRequestParams models.GetLoanRequestParams) models.GetLoanReturn {
	getLoanReturn := models.GetLoanReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = getLoanRequestParams.Symbol

	strRequest := "/v1/margin/loan-orders"
	jsonLoanReturn := utils.ApiKeyGet(c.Config, mapParams, strRequest)
	_ = json.Unmarshal([]byte(jsonLoanReturn), &getLoanReturn)

	return getLoanReturn
}

// 申请
// placeRequestParams: 下单信息
// return: PlaceReturn对象
func (c *RestClient) PlaceLoan(loanRequestParams models.LoanRequestParams) models.LoanReturn {
	loanReturn := models.LoanReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = loanRequestParams.Symbol
	mapParams["currency"] = loanRequestParams.Currency
	mapParams["amount"] = loanRequestParams.Amount

	strRequest := "/v1/margin/orders"
	jsonLoanReturn := utils.ApiKeyPost(c.Config, mapParams, strRequest)
	_ = json.Unmarshal([]byte(jsonLoanReturn), &loanReturn)

	return loanReturn
}

// 申请撤销一个订单请求
// strOrderID: 订单ID
// return: PlaceReturn对象
func (c *RestClient) CancelLoan(strOrderID string) models.LoanReturn {
	loanReturn := models.LoanReturn{}

	strRequest := fmt.Sprintf("/v1/margin/orders/%s/repay", strOrderID)
	jsonLoanReturn := utils.ApiKeyPost(c.Config, make(map[string]string), strRequest)
	_ = json.Unmarshal([]byte(jsonLoanReturn), &loanReturn)

	return loanReturn
}
