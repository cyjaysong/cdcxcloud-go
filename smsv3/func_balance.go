package smsv3

import (
	"encoding/json"
)

// Balance 余额查询
func (c *Client) Balance() (resBody *BalanceResBody, err error) {
	data := map[string]any{"action": "balance", "account": c.Account}
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
