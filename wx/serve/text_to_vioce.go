package serve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var  (
	ApiKey="ZHRLdu2E8FgqvNdPVh4e8H99"
	SecretKey="ESceg2zGaGi5dEr07kZVmnl2Ge7OXKg3"
	TokenApi = fmt.Sprintf(`https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s`,ApiKey,SecretKey)
	Cuid = "52:54:00:2e:f6:9e"
	AudioApi = "https://tsn.baidu.com/text2audio"
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


// ResponseToken 获取百度token
type ResponseToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn  int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
	SessionKey string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}

// PubPostAudioByText 获取语音
func PubPostAudioByText(text string, t *ResponseToken)(ret []byte,err error)  {
	var(
		ra = &RequestAudio{
			Tex:text,
			Tok:t.AccessToken,
			Ctp:1,
			Lan:"zh",
			CUID:Cuid,
			Spd:5,
			Pit:6,
			Vol:8,
			Per:0,
			Aue:6,
		}
		client = &http.Client{}
		resp *http.Response
		)

	r,_:=json.Marshal(ra)
	req,_err:=http.NewRequest("POST",AudioApi,bytes.NewReader(r))
	if _err!=nil{
		err =_err
		return
	}
	if resp,err=client.Do(req);err!=nil{
		return
	}

	defer resp.Body.Close()

	if ret ,err =ioutil.ReadAll(resp.Body);err!=nil{
		return
	}

	return
}

// PubGetAudioByText 获取语音
func PubGetAudioByText(text string, t *ResponseToken)(ret []byte,err error)  {
	var(
		ra = &RequestAudio{
			Tex:text,
			Tok:t.AccessToken,
			Ctp:1,
			Lan:"zh",
			CUID:Cuid,
			Spd:5,
			Pit:6,
			Vol:8,
			Per:0,
			Aue:6,
		}
		resp *http.Response
	)

	if resp,err=http.Get(fmt.Sprintf(`%s?lan=zh&ctp=1&cuid=%s&tok=%s&tex=%s`,AudioApi,"123456789.0",
		ra.Tok,text));err!=nil{
		return
	}

	defer resp.Body.Close()

	if ret ,err =ioutil.ReadAll(resp.Body);err!=nil{
		return
	}

	return
}
// RequestAudio
type RequestAudio struct {
	Tex string `json:"tex"` // 文本
	Tok string `json:"tok"` // token
	CUID string `json:"cuid"` // 用户唯一标识
	Lan string `json:"lan"`  // zh
	Ctp int `json:"ctp"` // 客户端类型选择，web端填写固定值1
	Spd int `json:"spd"` // 语速，取值0-15，默认为5中语速
	Pit int `json:"pit"` // 音调，取值0-15，默认为5中语调
	Vol int `json:"vol"` // 音量，取值0-15，默认为5中音量
	Per int `json:"per"` // 发音人选择, 0为普通女声，1为普通男生，3为情感合成-度逍遥，4为情感合成-度丫丫，默认为普通女声
	Aue int `json:"aue"` // 3为mp3格式(默认)； 4为pcm-16k；5为pcm-8k；6为wav（内容同pcm-16k）; 注意aue=4或者6是语音识别要求的格式， // 但是音频内容不是语音识别要求的自然人发音，所以识别效果会受影响。
}