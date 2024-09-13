package smsv2

import (
	"encoding/json"
)

// Report 获取状态报告
func (c *Client) Report(size int) (resBody *ReportResBody, err error) {
	data := map[string]any{"action": "report", "account": c.Account, "password": c.Password, "size": size}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &ReportResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}

// ReportNotifyVerify 被动接收状态报告
func (c *Client) ReportNotifyVerify(reqBodyBytes []byte) (notifyBody *ReportNotifyBody, err error) {
	notifyBody = &ReportNotifyBody{}
	if err = json.Unmarshal(reqBodyBytes, notifyBody); err != nil {
		return nil, err
	}
	return
}
