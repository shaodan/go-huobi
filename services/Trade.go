package services

import (
	"encoding/json"
	"fmt"
	"github.com/shaodan/go-huobi/models"
	"github.com/shaodan/go-huobi/utils"
)

//------------------------------------------------------------------------------------------
// 交易API

// 下单
// placeRequestParams: 下单信息
// return: PlaceReturn对象
func (c *HuobiRestClient) Place(placeRequestParams models.PlaceRequestParams) models.PlaceReturn {
	placeReturn := models.PlaceReturn{}

	mapParams := make(map[string]string)
	mapParams["account-id"] = placeRequestParams.AccountID
	mapParams["amount"] = placeRequestParams.Amount
	if 0 < len(placeRequestParams.Price) {
		mapParams["price"] = placeRequestParams.Price
	}
	if 0 < len(placeRequestParams.Source) {
		mapParams["source"] = placeRequestParams.Source
	}
	mapParams["symbol"] = placeRequestParams.Symbol
	mapParams["type"] = placeRequestParams.Type

	strRequest := "/v1/order/orders/place"
	jsonPlaceReturn := utils.ApiKeyPost(c.Config, mapParams, strRequest)
	_ = json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn
}

// 申请撤销一个订单请求
// strOrderID: 订单ID
// return: PlaceReturn对象
func (c *HuobiRestClient) SubmitCancel(strOrderID string) models.PlaceReturn {
	placeReturn := models.PlaceReturn{}

	strRequest := fmt.Sprintf("/v1/order/orders/%s/submitcancel", strOrderID)
	jsonPlaceReturn := utils.ApiKeyPost(c.Config, make(map[string]string), strRequest)
	_ = json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn
}
