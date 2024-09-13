package smsv2

import (
	reqclient "github.com/imroc/req/v3"
	"time"
)

const baseUrl = "https://api.cdcxcloud.com/smsv2"

type Client struct {
	Account   string // 账号
	Password  string // 账号
	ExtNo     string // 接入码
	reqClient *reqclient.Client
}

func NewClient(account, password, extNo string) *Client {
	client := &Client{Account: account, Password: password, ExtNo: extNo}
	client.reqClient = reqclient.C().SetTimeout(time.Second * 10).SetCommonRetryCount(2)
	client.reqClient.SetUserAgent("")
	return client
}
