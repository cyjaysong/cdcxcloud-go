package smsv2

import (
	"encoding/json"
	"time"
)

// Statis 获取统计信息
func (c *Client) Statis(beginTime, endTime time.Time) (resBody *StatisResBody, err error) {
	data := map[string]any{"action": "statis", "account": c.Account, "password": c.Password,
		"beginTime": beginTime.Format("20060102"), "endTime": endTime.Format("20060102")}
	bodyBytes, err := c.post(data)
	if err != nil {
		return nil, err
	}
	resBody = &StatisResBody{}
	if err = json.Unmarshal(bodyBytes, resBody); err != nil {
		return nil, err
	}
	return
}
