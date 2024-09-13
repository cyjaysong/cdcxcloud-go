package smsv3

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

func (c *Client) getSign(signVal string) string {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(signVal))
	return hex.EncodeToString(sha256Hash.Sum(nil))
}

func (c *Client) getEncryptBody(data map[string]any) string {
	jsonBytes, _ := json.Marshal(data)
	return aesEncryptByECB(jsonBytes, c.EncryptKey)
}

func (c *Client) post(data map[string]any) (bodyBytes []byte, err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	dataStr := c.getEncryptBody(data)
	sign := c.getSign(c.Password + dataStr + timestamp)

	req := c.reqClient.R().SetBodyJsonMarshal(map[string]string{"data": dataStr})
	res, err := req.SetHeader("sign", sign).SetHeader("timestamp", timestamp).Post(baseUrl)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, errors.New(string(res.Bytes()))
	}
	var resBody struct {
		Data string `json:"data"`
	}
	if err = res.UnmarshalJson(&resBody); err != nil {
		return nil, err
	}
	resSign, resTimestamp := res.GetHeader("sign"), res.GetHeader("timestamp")
	if sign = c.getSign(c.Password + resBody.Data + resTimestamp); sign != resSign {
		return nil, errors.New("response sign invalid")
	}
	bodyBytes = aesDecryptByECB(resBody.Data, c.EncryptKey)
	return
}

// GetNotifyResponseParam 获取回调通知响应参数
func (c *Client) GetNotifyResponseParam() (header, body map[string]string) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	data := c.getEncryptBody(map[string]any{"result": 0})
	sign := c.getSign(c.Password + data + timestamp)
	header = map[string]string{"spid": c.Account, "timestamp": timestamp, "sign": sign}
	body = map[string]string{"data": data}
	return
}

func aesEncryptByECB(dataBytes []byte, key string) string {
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic("key长度必须是 16、24、32 其中一个")
	}
	keyByte := []byte(key)
	block, _ := aes.NewCipher(keyByte)
	blockSize := block.BlockSize()
	dataBytes = pkcs7Padding(dataBytes, blockSize)
	encryptResult := make([]byte, len(dataBytes))
	for bs, be := 0, blockSize; bs < len(dataBytes); bs, be = bs+blockSize, be+blockSize {
		block.Encrypt(encryptResult[bs:be], dataBytes[bs:be])
	}
	return base64.StdEncoding.EncodeToString(encryptResult)
}

func pkcs7Padding(originByte []byte, blockSize int) []byte {
	padding := blockSize - len(originByte)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(originByte, padText...)
}

func aesDecryptByECB(data, key string) []byte {
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic("key长度必须是 16、24、32 其中一个")
	}
	// 反解密码base64
	originByte, _ := base64.StdEncoding.DecodeString(data)
	keyByte := []byte(key)
	block, _ := aes.NewCipher(keyByte)
	blockSize := block.BlockSize()
	decrypted := make([]byte, len(originByte))
	for bs, be := 0, blockSize; bs < len(originByte); bs, be = bs+blockSize, be+blockSize {
		block.Decrypt(decrypted[bs:be], originByte[bs:be])
	}
	return pkcs7UnPadding(decrypted)
}

func pkcs7UnPadding(originDataByte []byte) []byte {
	length := len(originDataByte)
	unpadding := int(originDataByte[length-1])
	return originDataByte[:(length - unpadding)]
}
