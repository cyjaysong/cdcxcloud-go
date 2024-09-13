package smsv2

type SendResult struct {
	Mid    string `json:"mid"`
	Mobile string `json:"mobile"`
	Result int    `json:"result"`
}

type BatchSendSmsResBody struct {
	Status  int          `json:"status"`
	Balance int          `json:"balance"`
	List    []SendResult `json:"list"`
}

type P2pSendSmsResBody struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	List    []SendResult `json:"list"`
}

type SelectSmsResBody struct {
	Status  int `json:"status"`
	Balance int `json:"balance"` // 当前账户余额，单位厘
	List    []struct {
		Apmid        string `json:"apmid"`        // 提交时响应的消息ID
		ApSubmitTime string `json:"apSubmitTime"` // 提交时间
		Mobile       string `json:"mobile"`       // 电话号码
		Status       int    `json:"status"`       // 0:客户提交成功;1:客户提交失败;2:转发提交成功;3:转发提交失败;4:用户接收成功;5:用户接收失败
		Stat         string `json:"stat"`         // 状态码，当status=4或5时，此值才有意义
		DeliverTime  string `json:"deliverTime"`  // 报告时间，当status=4或5，此值才有意义
	} `json:"list"`
}
