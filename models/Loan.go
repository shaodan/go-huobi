package models

type GetLoanRequestParams struct {
	Symbol string `json:"symbol"` // 交易对
}

type GetLoanReturn struct {
	Status  string      `json:"status"`
	Data    []LoanOrder `json:"data"`
	ErrCode string      `json:"err-code"`
	ErrMsg  string      `json:"err-msg"`
}

type LoanOrder struct {
	ID              int64  `json:"id"`               // 订单号
	UserID          int64  `json:"user-id"`          // 用户ID
	AccountID       int64  `json:"account-id"`       // 账户ID
	Symbol          string `json:"symbol"`           // 交易对
	Currency        string `json:"currency"`         //币种
	LoanAmount      string `json:"loan-amount"`      // 借贷本金总额
	LoanBalance     string `json:"loan-balance"`     // 未还本金
	InterestRate    string `json:"interest-rate"`    // 利率
	InterestAmount  string `json:"interest-amount"`  // 利息总额
	InterestBalance string `json:"interest-balance"` // 未还利息
	CreatedAt       int64  `json:"created-at"`       // 借贷发起时间
	AccruedAt       int64  `json:"accrued-at"`       // 最近一次计息时间
	State           string `json:"state"`            // 订单状态	created 未放款，accrual 已放款，cleared 已还清，invalid 异常
}

type LoanRequestParams struct {
	Symbol   string `json:"symbol"`   // 交易对, btcusdt, bccbtc......
	Currency string `json:"currency"` // 借贷币种
	Amount   string `json:"amount"`   // 借贷数量
}

type LoanReturn struct {
	Status  string `json:"status"`
	Data    int    `json:"data"`
	ErrCode string `json:"err-code"`
	ErrMsg  string `json:"err-msg"`
}

type RepayRequestParams struct {
	Symbol   string `json:"symbol"`   // 交易对, btcusdt, bccbtc......
	Currency string `json:"currency"` // 借贷币种
	Amount   string `json:"amount"`   // 借贷数量
}
