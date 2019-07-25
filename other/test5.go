package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var (
		bits           int
		err            error
		pubKey, priKey []byte
		n              int
	)
	flag.IntVar(&bits, "b", 1024, "密钥长度，默认为1024位")
	flag.Parse()
	// 生成密钥对
	if err = generateRsaKey(bits); err != nil {
		panic(err)
	}
	if priKey, err = ioutil.ReadFile("private.pem"); err != nil {
		return
	}
	if pubKey, err = ioutil.ReadFile("public.pem"); err != nil {
		return
	}
	// 获取长度
	if n, err = getPrivateKeyLen(priKey); err == nil {
		fmt.Printf("private key len is %d\n", n)
	} else {
		panic(err)
	}
	// 获取长度
	if n, err = getPublicKeyLen(pubKey); err == nil {
		fmt.Printf("public key len is %d\n", n)
	} else {
		panic(err)
	}

}

// generateRsaKey 密钥
func generateRsaKey(bits int) (err error) {
	var (
		file   *os.File
		priKey *rsa.PrivateKey
		pubKey *rsa.PublicKey
	)

	// 获取私钥
	if priKey, err = rsa.GenerateKey(rand.Reader, bits); err != nil {
		return
	}
	// 将私钥转换为ASN.1 DER编码形式。
	derStream := x509.MarshalPKCS1PrivateKey(priKey)
	// 私钥参数
	block := &pem.Block{
		Type:    "privateKey",
		Headers: map[string]string{"test": "testKey"},
		Bytes:   derStream,
	}
	// 创建private.key文件
	if file, err = os.Create("private.key"); err != nil {
		return
	}
	// 在文件里面写入数据
	if err = pem.Encode(file, block); err != nil {
		return
	}
	// 获取公钥
	pubKey = &priKey.PublicKey
	if _derPkix, _err := x509.MarshalPKIXPublicKey(pubKey); _err != nil {
		err = _err
		return
	} else {
		block = &pem.Block{
			Type:    "publicKey",
			Headers: map[string]string{"test": "testKey"},
			Bytes:   _derPkix,
		}
	}
	// 创建公钥文件
	if file, err = os.Create("public.key"); err != nil {
		return
	}
	// 写入数据
	if err = pem.Encode(file, block); err != nil {
		return
	}

	return
}

// getPublicKeyLen 获取公钥长度
func getPublicKeyLen(public []byte) (ret int, err error) {
	var (
		block *pem.Block
		data  interface{}
	)
	// 解码
	block, _ = pem.Decode(public)
	// 解析出数据
	if data, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		return
	}
	// 分析数据
	switch v := data.(type) {
	case *rsa.PublicKey:
		ret = v.N.BitLen()

	default:
		err = errors.New("data type is not rsa.PublicKey")
	}

	return
}

// getPrivateKeyLen 获取私钥长度
func getPrivateKeyLen(private []byte) (ret int, err error) {
	block, _ := pem.Decode(private)
	var (
		data *rsa.PrivateKey
	)
	if data, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return
	}

	ret = data.N.BitLen()
	return
}
