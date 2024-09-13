package smsv3_test

import "github.com/cyjaysong/cdcxcloud-go/smsv3"

var client *smsv3.Client

func init() {
	client = smsv3.NewClient("", "", "", "")
}
