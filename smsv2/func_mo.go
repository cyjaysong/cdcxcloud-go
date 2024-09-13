package smsv2

import (
	"encoding/json"
)

// Mo 获取手机上行回复
func (c *Client) Mo(size int) (resBody *MoResBody, err error) {
	data := map[string]any{"action": "mo", "account": c.Account, "password": c.Password, "size": size}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &MoResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}

// MoNotifyVerify 被动接收上行
func (c *Client) MoNotifyVerify(reqBodyBytes []byte) (notifyBody *MoNotifyBody, err error) {
	notifyBody = &MoNotifyBody{}
	if err = json.Unmarshal(reqBodyBytes, notifyBody); err != nil {
		return nil, err
	}
	return
}
