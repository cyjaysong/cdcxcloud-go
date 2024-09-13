package smsv3

import (
	"encoding/json"
	"errors"
	"time"
)

// TemplateP2pSend 模版点对点发送消息
func (c *Client) TemplateP2pSend(sendParams []TemplateP2pSendParam, templateId, label string, atTime *time.Time) (resBody *TemplateP2pSendResBody, err error) {
	variable := make([]map[string]string, len(sendParams))
	for i, param := range sendParams {
		param.Param["mobile"] = param.Mobile
		variable[i] = param.Param
	}
	data := map[string]any{"action": "templatep2p", "account": c.Account, "extno": c.ExtNo,
		"templateJson": map[string]any{"templateID": templateId, "variable": variable}}
	if len(label) > 0 {
		data["label"] = label
	}
	if atTime != nil {
		data["atTime"] = atTime.Format("2006-01-02 15:04:05")
	}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &TemplateP2pSendResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}

// TemplateAdd 模板添加
func (c *Client) TemplateAdd(templates []TemplateItem) (resBody *TemplateAddResBody, err error) {
	data := map[string]any{"action": "templateAdd", "account": c.Account, "templateJson": templates}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &TemplateAddResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}

// TemplateDelete 模板删除
func (c *Client) TemplateDelete(templateIds []int64) (resBody *TemplateDeleteResBody, err error) {
	data := map[string]any{"action": "templateDelete", "account": c.Account,
		"templateJson": map[string]any{"templateID": templateIds}}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &TemplateDeleteResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}

// TemplateSelect 模板查询
func (c *Client) TemplateSelect(templateID int64) (resBody *TemplateSelectResBody, err error) {
	data := map[string]any{"action": "templateSelect", "account": c.Account,
		"templateJson": map[string]any{"templateID": templateID}}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &TemplateSelectResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}

// TemplateNotifyVerify 模版审核通知
func (c *Client) TemplateNotifyVerify(spid, sign, timestamp string, reqBodyBytes []byte) (notifyBody *TemplateNotifyBody, err error) {
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
	notifyBody = &TemplateNotifyBody{}
	bodyBytes := aesDecryptByECB(reqBody.Data, c.EncryptKey)
	if err = json.Unmarshal(bodyBytes, notifyBody); err != nil {
		return nil, err
	}
	return
}
