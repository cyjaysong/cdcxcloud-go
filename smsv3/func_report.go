package smsv3

import (
	"encoding/json"
	"errors"
)

// Report 获取状态报告
func (c *Client) Report(size int) (resBody *ReportResBody, err error) {
	data := map[string]any{"action": "report", "account": c.Account, "size": size}
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
func (c *Client) ReportNotifyVerify(spid, sign, timestamp string, reqBodyBytes []byte) (notifyBody *ReportNotifyBody, err error) {
	if c.Account != spid {
		return nil, errors.New("spid invalid ")
	}
	var reqBody struct {
		Data string `json:"data"`
	}
	if err = json.Unmarshal(reqBodyBytes, &reqBody); err != nil {
		return nil, err
	}
	if mySign := c.getSign(c.Password + reqBody.Data + timestamp); mySign != sign {
		return nil, errors.New("sign invalid")
	}
	notifyBody = &ReportNotifyBody{}
	bodyBytes := aesDecryptByECB(reqBody.Data, c.EncryptKey)
	if err = json.Unmarshal(bodyBytes, notifyBody); err != nil {
		return nil, err
	}
	return
}
