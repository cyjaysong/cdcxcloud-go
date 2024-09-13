package smsv2

type PatternAddResBody struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	List    []struct {
		Spid      string `json:"spid"`
		PatternID int    `json:"patternID"`
		Pattern   string `json:"pattern"`
		Status    int    `json:"status"`
	} `json:"list"`
}
