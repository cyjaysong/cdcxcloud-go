package smsv3

type ReportItem struct {
	Mid        string `json:"mid"`        // 消息ID(与提交时响应的消息ID匹配)
	Spid       string `json:"spid"`       // 账户ID
	AccessCode string `json:"accessCode"` // 下发号码
	Mobile     string `json:"mobile"`     // 手机号码
	Stat       string `json:"stat"`       // 状态报告代码，参见STAT状态报告代码表
	Time       string `json:"time"`       // 报告时间，格式：2022-12-05 18:00:00
	Label      string `json:"label"`      // 下发短信提交带的label
}

type ReportResBody struct {
	Status  int          `json:"status"`  // 请求结果，具体参见STATUS错误代码表
	Message string       `json:"message"` // 请求结果说明，具体见STATUS错误代码表
	List    []ReportItem `json:"list"`
}

type ReportNotifyBody []ReportItem
