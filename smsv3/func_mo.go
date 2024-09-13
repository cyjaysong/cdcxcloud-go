package smsv3

import (
	"encoding/json"
	"errors"
)

// Mo 获取手机上行
func (c *Client) Mo(size int) (resBody *MoResBody, err error) {
	data := map[string]any{"action": "mo", "account": c.Account, "size": size}
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
func (c *Client) MoNotifyVerify(spid, sign, timestamp string, reqBodyBytes []byte) (notifyBody *MoNotifyBody, err error) {
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
	notifyBody = &MoNotifyBody{}
	bodyBytes := aesDecryptByECB(reqBody.Data, c.EncryptKey)
	if err = json.Unmarshal(bodyBytes, notifyBody); err != nil {
		return nil, err
	}
	return
}
