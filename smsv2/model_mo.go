package smsv2

type MoItem struct {
	Mid        string `json:"mid"`        // 上行消息ID
	Spid       string `json:"spid"`       // 账户ID
	AccessCode string `json:"accessCode"` // 接收号码，即SP服务号（106XXXXXX）
	Mobile     string `json:"mobile"`     // 手机号码
	Content    string `json:"content"`    // 上行短信内容
	Time       string `json:"time"`       // 上行时间，格式：2022-12-05 18:00:00
}

type MoResBody struct {
	Status  int      `json:"status"`  // 请求结果，具体参见STATUS错误代码表
	Message string   `json:"message"` // 请求结果说明，具体见STATUS错误代码表
	List    []MoItem `json:"list"`
}

type MoNotifyBody []ReportItem
