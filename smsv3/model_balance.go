package smsv3

type BalanceResBody struct {
	Status     int    `json:"status"`     // 请求结果，具体参见STATUS错误代码表
	Message    string `json:"message"`    // 请求结果说明，具体见STATUS错误代码表
	ChargeType string `json:"chargeType"` // 付费类型：POSTCHARGE=后付费；PRECHARGE=预付费
	Balance    int    `json:"balance"`    // 当前账户余额，单位厘
	Number     int    `json:"number"`     // 账户剩余条数
}
