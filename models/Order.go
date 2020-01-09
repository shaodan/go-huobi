package models

import "encoding/json"

type Order struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Symbol       string `json:"symbol"`
	Type         string `json:"type"` // sell-market, buy-limit...
	Amount       string `json:"amount"`
	State        string `json:"state"` // filled, partial-canceled, canceled
	FilledAmount string `json:"filled-amount"`
	Source       string `json:"source"`
	CreatedAt    int64  `json:"created-at"`
	AccountID    int64  `json:"account-id"`
}

type OrderTypo struct {
	Order
	FieldAmount string `json:"field-amount"`
}

func DecodeOrder(raw []byte) (*Order, error) {
	var order = &Order{}
	if err := json.Unmarshal(raw, order); err != nil {
		return nil, err
	}
	return order, nil
}

func DecodeOrderTypo(raw []byte) (*Order, error) {
	var order = &OrderTypo{}
	if err := json.Unmarshal(raw, order); err != nil {
		return nil, err
	}
	order.FilledAmount = order.FieldAmount
	return &order.Order, nil
}
