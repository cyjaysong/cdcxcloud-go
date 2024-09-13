package smsv3

import (
	reqclient "github.com/imroc/req/v3"
	"time"
)

const baseUrl = "http://api.cdcxcloud.com:7862/smsv3"

type Client struct {
	Account    string // 账号
	Password   string // 账号
	ExtNo      string // 接入码
	EncryptKey string // 加解密秘钥
	reqClient  *reqclient.Client
}

func NewClient(account, password, extNo, encryptKey string) *Client {
	client := &Client{Account: account, Password: password, ExtNo: extNo, EncryptKey: encryptKey}
	client.reqClient = reqclient.C().SetTimeout(time.Second * 10).SetCommonRetryCount(2)
	client.reqClient.SetCommonHeader("spid", account)
	client.reqClient.SetUserAgent("")
	return client
}
