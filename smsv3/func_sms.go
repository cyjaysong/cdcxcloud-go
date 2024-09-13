package smsv3

import (
	"encoding/json"
	"strings"
	"time"
)

// BatchSendSms 批量发送消息
// [mobiles]批量发送的电话号码(去重); [content]短信内容; [label]用于报告返回的标记内容; [atTime]指定时间发送,若为空或者定时在5分钟内则立即发送
func (c *Client) BatchSendSms(mobiles []string, content, label string, atTime *time.Time) (resBody *BatchSendSmsResBody, err error) {
	data := map[string]any{"action": "send", "account": c.Account, "extno": c.ExtNo,
		"mobile": strings.Join(mobiles, ","), "content": content}
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
	resBody = &BatchSendSmsResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}

// P2pSendSms 点对点发送消息
// [mobileContentKv]号码内容键值对; [label]用于报告返回的标记内容; [atTime]指定时间发送,若为空或者定时在5分钟内则立即发送
func (c *Client) P2pSendSms(mobileContentKv map[string]string, label string, atTime *time.Time) (resBody *P2pSendSmsResBody, err error) {
	data := map[string]any{"action": "p2p", "account": c.Account, "extno": c.ExtNo,
		"mobileContentKvp": mobileContentKv}
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
	resBody = &P2pSendSmsResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}
