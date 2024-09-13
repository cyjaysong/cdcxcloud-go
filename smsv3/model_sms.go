package smsv3

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
