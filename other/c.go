package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	var (
		pub = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrd5QrtGfFO7tSikfpz9qx7cmR
QIVnNio+exGE0HuqEbhQCC/5YhTjELc4k7uYhkKto6ucfZy3TUf6dstqLWaL/VWA
vlvX4CC5m49f20n67T+DR54xFAmGCzVf0a+RCnh+uB5cZi1FvgmB/1MoRWKvbiwM
TH5R4y9sUhu3SQ43HQIDAQAB`
		pri = `MIICXAIBAAKBgQCrd5QrtGfFO7tSikfpz9qx7cmRQIVnNio+exGE0HuqEbhQCC/5
YhTjELc4k7uYhkKto6ucfZy3TUf6dstqLWaL/VWAvlvX4CC5m49f20n67T+DR54x
FAmGCzVf0a+RCnh+uB5cZi1FvgmB/1MoRWKvbiwMTH5R4y9sUhu3SQ43HQIDAQAB
AoGAJuT/FVLc3x6HhVecrGrbvtSjjnFGUX014+pitO/dvVw7pNvWlgkrl74o8YqB
WT3LTjv8J4lOzT2YgDYGOlWGFwf7yCOgZRpt8xlSu3ZydRPm0sxEvWgv0JYny9m+
EhZ0FqeqH7QWlHdkO8HqvJQva8u7E2y/O8BbT/YMiu7qpMECQQDcEwLvjHBBXq+M
c7sFonNojzUMBYL7B8xY1P/Ot0Yo+7BEwsdEtSIGj5akZfuFr9QZgruBtezFNnZM
ug+KjzAZAkEAx3VBX8NfTkBO4nmwUUyLXplApOzKZGlYVSNYcUkptIEpT2wgBxUl
ghSknFsaGbFpGsBEkZCgXNftNxI3peDPpQJAY44aIuGmGnxJ78Ce1yKxEJjQB3sq
0IKrl3frrMjN7VZGXCS83kEOfmdQX1hfGw/6Y/v29Oumi2RiybzVsPmraQJAA5/u
4zWiusJSbK03dhLFCaARW63t86sybsGors5ckqoyPP5DCr3oo5eKckj5jXP67ACI
fni5YVaPOgv7tOkD/QJBAJbS0DoBYfYahYT7Y85DpJMHYDfVuVm8fzupapPa6KRx
ODoXHa3JzkhB5r6/W+l8T6JVeBWo74+OC6UFfqIl1Is=
`
	)

	s, err := RsaEncryptWithSha1Base64("哈哈", pub)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Println(RsaDecryptWithSha1Base64("RulcTZ2QRms94LvDRZ8r2g61FKu+u2dl06IeMgAIDP4bFY+gvZd03lPCkd33juB5kPv0SSsdOyDRf5NxbIPkM7IAALZSuCsu9kdJuaFb74l1+T98NC5R3CEVAtOUD7GyFLQGev/ixjaEezBRAjgTJt7BPt5HzSzJ5t1VJW/9NRU=", pri))
}

func RsaEncryptWithSha1Base64(originalData, publicKey string) (ret string, err error) {
	key, _ := base64.StdEncoding.DecodeString(publicKey)
	pubKey, _ := x509.ParsePKIXPublicKey(key)
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(originalData))
	return base64.StdEncoding.EncodeToString(encryptedData), err
}

func RsaDecryptWithSha1Base64(encryptedData, privateKey string) (string, error) {
	encryptedDecodeBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	key, _ := base64.StdEncoding.DecodeString(privateKey)
	prvKey, _ := x509.ParsePKCS1PrivateKey(key)
	originalData, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey, encryptedDecodeBytes)
	return string(originalData), err
}

func RsaSignWithSha1Hex(data string, prvKey string) (string, error) {
	keyByts, err := hex.DecodeString(prvKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(keyByts)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		return "", err
	}
	h := sha1.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	out := hex.EncodeToString(signature)
	return out, nil
}

func RsaVerySignWithSha1Base64(originalData, signData, pubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
}
