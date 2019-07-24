package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"github.com/srlemon/note"
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
	if err = genRsaKey(bits); err != nil {
		panic(err)
	}
	if priKey, err = ioutil.ReadFile("private.pem"); err != nil {
		return
	}
	if pubKey, err = ioutil.ReadFile("public.pem"); err != nil {
		return
	}
	if n, err = getPriKeyLen(priKey); err == nil {
		fmt.Printf("private key len is %d\n", n)
	} else {
		panic(err)
	}
	if n, err = getPubKeyLen(pubKey); err == nil {
		fmt.Printf("public key len is %d\n", n)
	} else {
		panic(err)
	}

}

// genRsaKey 生成密钥
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
		Type:  "私钥", // 私钥type
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
		Type:  "公钥",
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

// getPubKeyLen 获取公钥长度
func getPubKeyLen(pubkey []byte) (ret int, err error) {
	if pubkey == nil {
		err = note.ErrKeyIsNull
		return
	}
	var (
		block        *pem.Block
		pubInterface interface{}
	)
	if block, _ = pem.Decode(pubkey); block == nil {
		err = note.ErrPubKeyRsa
		return
	}

	if pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		return
	}
	pub := pubInterface.(*rsa.PublicKey)
	ret = pub.N.BitLen()
	return
}

// getPriKeyLen 获取私钥长度
func getPriKeyLen(priKey []byte) (ret int, err error) {
	if priKey == nil {
		err = note.ErrKeyIsNull
		return
	}
	block, _ := pem.Decode(priKey)
	if block == nil {
		err = note.ErrPriKerRsa
		return
	}
	var (
		pri *rsa.PrivateKey
	)
	if pri, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return
	}
	ret = pri.N.BitLen()
	return
}
