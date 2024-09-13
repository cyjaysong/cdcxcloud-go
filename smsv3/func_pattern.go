package smsv3

import (
	"encoding/json"
)

// PatternAdd 内容报备
func (c *Client) PatternAdd(pattern []string) (resBody *PatternAddResBody, err error) {
	data := map[string]any{"action": "patternAdd", "account": c.Account, "pattern": pattern}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &PatternAddResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}

// PatternSignAdd 签名报备
func (c *Client) PatternSignAdd(pattern []string) (resBody *PatternAddResBody, err error) {
	data := map[string]any{"action": "patternSignAdd", "account": c.Account, "pattern": pattern}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &PatternAddResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}
