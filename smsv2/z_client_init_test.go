package smsv2_test

import "github.com/cyjaysong/cdcxcloud-go/smsv2"

var client *smsv2.Client

func init() {
	client = smsv2.NewClient("", "", "")
}
