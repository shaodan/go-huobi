package models

// 账户信息
type AccountsReturn struct {
	Status  string       `json:"status"` // 请求状态
	Data    []SubAccount `json:"data"`   // 用户数据
	ErrCode string       `json:"err-code"`
	ErrMsg  string       `json:"err-msg"`
}

// 子账户信息
type SubAccount struct {
	ID      int64  `json:"id"`      // Account ID
	Type    string `json:"type"`    // 账户类型, spot：现货账户， margin：杠杆账户，otc：OTC 账户，point：点卡账户
	State   string `json:"state"`   // 账户状态, working: 正常, lock: 账户被锁定
	SubType string `json:"subtype"` // 子账户币种信息
	UserID  int64  `json:"user-id"` // 用户ID
}

// 账户余额
type Balance struct {
	ID     int64        `json:"id"`    // 账户ID
	State  string       `json:"state"` // 账户状态, working: 正常, lock: 账户被锁定
	Type   string       `json:"type"`  // 账户类型, spot: 现货账户
	List   []SubBalance `json:"list"`  // 子账户数组
	UserID int64        `json:"user-id"`
}

// 子账户余额
type SubBalance struct {
	Currency string `json:"currency"` // 币种
	Balance  string `json:"balance"`  // 结余
	Type     string `json:"type"`     // 类型, trade: 交易余额, frozen: 冻结余额
}

type BalanceReturn struct {
	Status  string  `json:"status"` // 请求状态
	Data    Balance `json:"data"`   // 账户余额
	ErrCode string  `json:"err-code"`
	ErrMsg  string  `json:"err-msg"`
}

type TransferType string

const (
	MasterTransferIn       = TransferType("master-transfer-in")        // 子账号划转给母账户虚拟币
	MasterTransferOut      = TransferType("master-transfer-out")       // 母账户划转给子账号虚拟币
	MasterPointTransferIn  = TransferType("master-point-transfer-in")  // 子账号划转给母账户点卡
	MasterPointTransferOut = TransferType("master-point-transfer-out") // 母账户划转给子账号点卡
)

type TransferRequestParams struct {
	SubUID   string       // 子账号uid
	Currency string       // 币种
	Amount   string       // 划转金额
	Type     TransferType // 划转类型
}

type TransferReturn struct {
	Status string `json:"status"` // OK / Error
	Data   int64  `json:"data"`   // 划转订单id
	// Error code
	// account-transfer-balance-insufficient-error  账户余额不足
	// base-operation-forbidden  禁止操作（母子账号关系错误时报）
}

type WithdrawRequestParams struct {
	Address  string // 提现钱包地址
	Currency string // 币种
	Amount   string // 提现数量
}

type WithdrawReturn struct {
	Data int64 `json:"data"` // 提现ID
}
