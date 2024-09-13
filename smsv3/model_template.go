package smsv3

type TemplateP2pSendParam struct {
	Mobile string            // 电话号码
	Param  map[string]string `json:"param"` // 变量参数
}

type TemplateP2pSendResBody struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	List    []SendResult `json:"list"`
}

type TemplateItem struct {
	Spid       string `json:"spid,omitempty"`       // 账户ID
	TemplateID int    `json:"templateID,omitempty"` // 模版ID
	Name       string `json:"name"`                 // 模板名
	TimePeriod string `json:"timePeriod"`           // 模板可用的时间；可为空，为空则表示不限制时间段，格式08:00-18:00
	Content    string `json:"content"`              // 模板内容；例如: 模板1[(B)]内容[(C)]  其中[(*)]为变量固定格式，*为任意可见字母
	Status     int    `json:"status,omitempty"`     // 模板状态；0:待审核；1:通过，2:拒绝，4:停用
}

type TemplateAddResBody struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	List    []TemplateItem `json:"list"` // 一次最多1000个模板；
}

type TemplateDeleteResBody struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type TemplateSelectResBody struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	List    []struct {
		Spid       string `json:"spid"`
		TemplateID int    `json:"templateID"`
		Name       string `json:"name"`
		TimePeriod string `json:"timePeriod"`
		Content    string `json:"content"`
		Status     int    `json:"status"`
	} `json:"list"`
}

type TemplateNotifyBody []TemplateItem
