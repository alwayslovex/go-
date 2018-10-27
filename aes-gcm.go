package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

//开发过程中,遇到了,使用c++ openssl库来进行aes-gcm 128位进行加密.然后将加密后的内容,传输给go服务.
//在网上查了下资料发现,使用go的解密库解密,有些信息不对称.
//其实go中的aes加密,后的内容,是加密内容和mac拼接在一起的.所以解密的时候也要进行拼接才能解密
func DecryptContent(encryKey, Content, Mac, Iv string) (string, error) {
	//开始解密
	block, err := aes.NewCipher([]byte(encryKey))
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	content, err := hex.DecodeString(Content) //由于将密文转换成了16进制字符.所以这里要转回去.
	if err != nil {
		return "", err
	}

	mac, err := hex.DecodeString(Mac)
	if err != nil {
		return "", err
	}

	contentAndMAc := append(content, mac...) //这里进行了拼接.
	decodeContent, err := aesGcm.Open(nil, []byte(Iv), contentAndMAc, nil)
	return string(decodeContent), err
}

func main() {
	DecryptContent("密钥", "加密后的内容", "校验码", "初始化向量")
}
