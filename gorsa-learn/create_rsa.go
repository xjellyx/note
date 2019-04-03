package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"os"
)

func main() {
	var (
		bits int
		err  error
	)
	flag.IntVar(&bits, "b", 1024, "密钥长度，默认为1024位")
	flag.Parse()
	if err = genRsaKey(bits); err != nil {
		panic(err)
	}
}

func genRsaKey(bits int) (err error) {
	var (
		file       *os.File
		privateKey *rsa.PrivateKey
		derPkix    []byte
	)
	// 获取密钥
	if privateKey, err = rsa.GenerateKey(rand.Reader, bits); err != nil {
		return
	}
	// 解析
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "LnFen私钥", // 私钥type
		Bytes: derStream,
	}
	if file, err = os.Create("private.pem"); err != nil {
		return
	}
	// 写入文件
	if err = pem.Encode(file, block); err != nil {
		return
	}
	publicKey := &privateKey.PublicKey
	if derPkix, err = x509.MarshalPKIXPublicKey(publicKey); err != nil {
		return
	}
	block = &pem.Block{
		Type:  "LnFen公钥",
		Bytes: derPkix,
	}
	if file, err = os.Create("public.pem"); err != nil {
		return
	}
	if err = pem.Encode(file, block); err != nil {
		return
	}
	return
}
