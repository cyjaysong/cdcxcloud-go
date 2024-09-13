package smsv2

type StatisResBody struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	List    []struct {
		StatisTime  string  `json:"statisTime"`
		Total       int     `json:"total"`
		Success     int     `json:"success"`
		Unknown     int     `json:"unknown"`
		Fail        int     `json:"fail"`
		SuccessRate float64 `json:"successRate"`
	} `json:"list"`
}
