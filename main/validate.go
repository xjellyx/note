package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
)

func getEcdsaKey(length int) (*ecdsa.PrivateKey, ecdsa.PublicKey, error) {
	var err error
	var prk *ecdsa.PrivateKey
	var puk ecdsa.PublicKey
	var curve elliptic.Curve
	switch length {
	case 1:
		curve = elliptic.P224()
	case 2:
		curve = elliptic.P256()
	case 3:
		curve = elliptic.P384()
	case 4:
		curve = elliptic.P521()
	default:
		err = errors.New("输入的签名等级错误！")
	}
	prk, err = ecdsa.GenerateKey(curve, rand.Reader) //通过 "crypto/rand" 模块产生的随机数生成私钥
	if err != nil {
		return prk, puk, err
	}
	puk = prk.PublicKey
	fmt.Println(prk.Sign())
	return prk, puk, err
}

func main() {
	getEcdsaKey(2)
}
