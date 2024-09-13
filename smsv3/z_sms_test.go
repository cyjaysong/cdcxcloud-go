package smsv3_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_P2pSendSms(t *testing.T) {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
	content := fmt.Sprintf("【哇哇哈哈】您的验证码为：%s，5分钟内有效，不要告诉任何人哦！如非本人操作，请忽略本短信。", code)
	mobileContentKv := map[string]string{"15066668888": content}
	sms, err := client.P2pSendSms(mobileContentKv, "", nil)
	fmt.Println(err)
	fmt.Printf("%+v\n", sms)
}
