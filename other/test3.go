package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var  (
	ApiKey="ZHRLdu2E8FgqvNdPVh4e8H99"
	SecretKey="ESceg2zGaGi5dEr07kZVmnl2Ge7OXKg3"
	TokenApi = fmt.Sprintf(`https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s`,ApiKey,SecretKey)

)

// PubGetToken 获取token
func PubGetToken()(ret ResponseToken,err error)  {
	var(
		data ResponseToken
		body []byte
		resp *http.Response

	)
	// 获取token
	if resp,err=http.Get(TokenApi);err!=nil{
		return
	}
	defer resp.Body.Close()
	if body,err=ioutil.ReadAll(resp.Body);err!=nil{
		return
	}
	if err = json.Unmarshal(body,&data);err!=nil{
		return
	}
	ret = data
	return
}

// 获取百度token
type ResponseToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn  int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
	SessionKey string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}

func main()  {
	fmt.Println(PubGetToken())
	var(
		client = &http.Client{}
		uri="https://wx2.qq.com/cgi-bin/mmwebwx-bin/webwxstatusnotify?"
	)
	km:=url.Values{}
	km.Add("pass_ticket","GISOyg2CCYGPvTwHrjo8Q5EFAe6xuVCXlz1u5HXT3I9Nc%2FoD%2FpoXmmXuZXPwqqUH")
	km.Add("lang","en_US")
	uri=uri+km.Encode()
	rand.Seed(time.Now().UnixNano())
	data := rand.Int63n(1000)
	js:= map[string]interface{}{
		"BaseRequest": map[string]interface{}{
			"Uin":1820469237,
			"Sid":"SKI8d6hTWF326hEM",
			"SKey":"@crypt_c9495c9f_4d97dc268b673bc8a6edff363e6130aa",
			"DeviceID":"e716353354772627",

		},
		"Code":1,
		"FromUserName":"@3db498e88028bcf9d19d686835d8d18c5a985638adc9c61988f9af4cdceace93",
		"ToUserName":"@47d9761265ad3724b1bf6cd1d716c96297b28339ea0a16f651e7795ab4c87be2",
		"ClientMsgId": int(time.Now().Unix()*1000 + data),
	}
	headers:= map[string]string{"Content-Type": "application/json; charset=UTF-8"}
	body,_:=json.Marshal(js)
	req,err:=http.NewRequest("POST",uri,bytes.NewReader(body))
	if err!=nil{
		panic(err)
	}

	req.Header.Set("User-Agent","Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

	defer req.Body.Close()
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp,err:=client.Do(req)
	if err!=nil{
		panic(err)
	}
	ret,_:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(ret))
}