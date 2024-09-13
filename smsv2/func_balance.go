package smsv2

import (
	"encoding/json"
)

// Balance 余额查询
func (c *Client) Balance() (resBody *BalanceResBody, err error) {
	data := map[string]any{"action": "balance", "account": c.Account, "password": c.Password}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &BalanceResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}

// BalanceNumber 余额条数查询
func (c *Client) BalanceNumber() (resBody *BalanceNumberResBody, err error) {
	data := map[string]any{"action": "balance", "account": c.Account, "password": c.Password}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &BalanceNumberResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}
