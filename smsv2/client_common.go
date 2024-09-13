package smsv2

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

func (c *Client) post(data map[string]any) (bodyBytes []byte, err error) {
	res, err := c.reqClient.R().SetBodyJsonMarshal(data).Post(baseUrl)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(res.Bytes()))
	}
	bodyBytes = res.Bytes()
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
