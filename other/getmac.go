package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

// Computer 计算机属性
type Computer struct {
	MacAddr string `json:"mac_addr"`
	IPAddr  string `json:"ip_addr"`
}

func main() {
	var (
		publicKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrd5QrtGfFO7tSikfpz9qx7cmR
QIVnNio+exGE0HuqEbhQCC/5YhTjELc4k7uYhkKto6ucfZy3TUf6dstqLWaL/VWA
vlvX4CC5m49f20n67T+DR54xFAmGCzVf0a+RCnh+uB5cZi1FvgmB/1MoRWKvbiwM
TH5R4y9sUhu3SQ43HQIDAQAB`
		addrs []string
		data  string
		resp  *http.Response
		body  []byte
		c     = new(Computer)
		err   error
	)
	defer func() {
		if err != nil {
			fmt.Println("拒绝非法机器执行程序")
			return
		}
	}()
	if addrs, err = GetMacAddr(); err != nil {
		return
	}
	if c.IPAddr, err = GetPublicIPAddr(); err != nil {
		return
	}
	c.MacAddr = addrs[0]
	origin, _ := json.Marshal(c)
	if data, err = EncryptMainRun(string(origin), publicKey); err != nil {
		return
	}

	if resp, err = http.PostForm("https://lisence.cjd88.cn/api/lisence/checkServerLisence", url.Values{
		"encrypt_key": {data},
	}); err != nil {
		return
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	var d struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Sign string `json:"sign"`
		}
	}
	fmt.Println(string(body))
	if err = json.Unmarshal(body, &d); err != nil {
		return
	}
	fmt.Println(string(body), "aaaaaaaaaa")
	if err = RSASignValid(string(origin), d.Data.Sign, publicKey); err != nil {
		return
	}
	fmt.Println("验证通过")
}

// GetMacAddr 获取mac地址
func GetMacAddr() (ret []string, err error) {
	var (
		netInterfaces []net.Interface
	)
	// 获取mac地址
	if netInterfaces, err = net.Interfaces(); err != nil {
		return
	}
	for _, v := range netInterfaces {
		macAddr := v.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		ret = append(ret, macAddr)
	}
	if len(ret) == 0 {
		err = errors.New("mac addr  no one")
		return
	}
	return
}

// EncryptMainRun main函数执行前加密
func EncryptMainRun(original, pubKey string) (ret string, err error) {
	var (
		data        []byte
		publicKey   interface{}
		encryptData []byte
	)
	if data, err = base64.StdEncoding.DecodeString(pubKey); err != nil {
		return
	}
	if publicKey, err = x509.ParsePKIXPublicKey(data); err != nil {
		return
	}
	if encryptData, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), []byte(original)); err != nil {
		return
	}

	//
	ret = base64.StdEncoding.EncodeToString(encryptData)
	return
}

// RSASignValid 验证签名
func RSASignValid(original, sign, public string) (err error) {
	var (
		decodeSign   []byte
		decodePublic []byte
		pub          interface{}
	)

	//
	if decodeSign, err = base64.StdEncoding.DecodeString(sign); err != nil {
		return
	}
	if decodePublic, err = base64.StdEncoding.DecodeString(public); err != nil {
		return
	}
	if pub, err = x509.ParsePKIXPublicKey(decodePublic); err != nil {
		return
	}
	hash := sha1.New()
	hash.Write([]byte(original))

	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), decodeSign)
}

// GetIPAddr 获取ip地址
func GetPublicIPAddr() (ret string, err error) {
	var (
		resp *http.Response
		body []byte
	)
	if resp, err = http.Get("http://ip.taobao.com/service/getIpInfo2.php?ip=myip"); err != nil {
		err = errors.New(` get public ip address fail`)
		return
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	var d struct {
		Data struct {
			IP string `json:"ip"`
		}
	}
	_ = json.Unmarshal(body, &d)
	if len(d.Data.IP) == 0 {
		err = errors.New(` get public ip address fail`)
		return
	}

	//
	ret = d.Data.IP
	return
}
