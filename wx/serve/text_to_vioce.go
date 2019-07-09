package serve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"time"
)

var  (
	ApiKeyBaidu="ZHRLdu2E8FgqvNdPVh4e8H99"
	SecretKeyBaidu="ESceg2zGaGi5dEr07kZVmnl2Ge7OXKg3"
	TokenApi = fmt.Sprintf(`https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s`,ApiKeyBaidu,SecretKeyBaidu)
	AudioApiBaidu = "https://tsn.baidu.com/text2audio"
	AudioApiTenxun = "https://aai.tencentcloudapi.com/?"
	SecretKeyTenxun = "AKIDV7DNUB5lZDeB7ru4mvSfzEibJ2JRszky"
)



// PubGetToken 获取token
func PubGetToken()(ret ResponseTokenBaidu,err error)  {
	var(
		data ResponseTokenBaidu
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


// ResponseTokenBaidu 获取百度token
type ResponseTokenBaidu struct {
	AccessToken string `json:"access_token"`
	ExpiresIn  int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
	SessionKey string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}
// RequestAudio
type RequestAudioBaidu struct {
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

// NewAudioBaidu
func NewAutioBaidu(s *RequestAudioBaidu)(d *RequestAudioBaidu,err error)  {
	if s!=nil{
		d = s
	}else {
		d = new(RequestAudioBaidu)
	}
	if err = d.InitDefault();err!=nil{
		return
	}
	return
}

// InitDefault
func (r *RequestAudioBaidu)InitDefault() (err error) {
	if r==nil{
		err = errors.New("requestAudio is nil ")
		return
	}
	if len(r.CUID)==0{
		r.CUID=uuid.NewV4().String()
	}
	if len(r.Lan)==0{
		r.Lan="zh"
	}
	if r.Ctp==0{
		r.Ctp=1
	}
	if r.Spd==0{
		r.Spd=5
	}
	if r.Pit==0{
		r.Pit = 5
	}
	if r.Vol==0{
		r.Vol=5
	}
	if r.Aue==0{
		r.Aue=6
	}
	return
}

// PubPostAudioByTextBaidu 获取语音
func PubPostAudioByTextBaidu(text string, t *ResponseTokenBaidu)(ret []byte,err error)  {
	var(
		client = &http.Client{}
		resp *http.Response
		rb *RequestAudioBaidu
		)

	if rb,err = NewAutioBaidu(nil);err!=nil{
		return
	}
	rb.Tex=text
	rb.CUID=uuid.NewV4().String()
	rb.Tok=t.AccessToken
	r,_:=json.Marshal(rb)
	req,_err:=http.NewRequest("POST",AudioApiBaidu,bytes.NewReader(r))
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

// PubGetAudioByTextBaidu 获取语音
func PubGetAudioByTextBaidu(text string, t *ResponseTokenBaidu)(ret []byte,err error)  {
	var(
		resp *http.Response
		rb *RequestAudioBaidu
	)

	if rb,err = NewAutioBaidu(nil);err!=nil{
		return
	}
	rb.CUID=uuid.NewV4().String()
	rb.Tok=t.AccessToken


	if resp,err=http.Get(fmt.Sprintf(`%s?lan=zh&ctp=1&cuid=%s&tok=%s&tex=%s`,AudioApiBaidu,rb.CUID,
		rb.Tok,text));err!=nil{
		return
	}

	defer resp.Body.Close()

	if ret ,err =ioutil.ReadAll(resp.Body);err!=nil{
		return
	}

	return
}

// RequestAudioTenxun
type RequestAudioTenxun struct {
	// 必填
	Action string `json:"action"` // 公共参数，本接口取值：TextToVoice
	Version string `json:"version"` // 公共参数，本接口取值：2018-05-22
	Region string `json:"region"` // 地域
	Text string `json:"text"`
	SessionId string `json:"sessionId"` // 请求session
	ModelType int `json:"modelType"` // 默认1
	// 选填
	Volume float32 `json:"volume"` // 音量大小[0-10],默认0
	Speed float32 `json:"speed"` // 语速, [-2,2],2代表0.6倍 -1代表0.8倍 0代表1.0倍（默认） 1代表1.2倍 2代表1.5倍
	ProjectId int `json:"projectId"` // 项目id,默认0
	VoiceType int `json:"voiceType"` // 音色 0-亲和女声(默认) 1-亲和男声 2-成熟男声 3-活力男声 4-温暖女声 5-情感女声 6-情感男声
	PrimaryLanguage int `json:"primaryLanguage"` // 语言类型, 1 中文,默认;2-英文
	SampleRate int `json:"sampleRate"` // 音频采样率 默认16k, 8000: 8k
	Codec string `json:"codec"` // wav默认; mp3
}

// ResponseAudioTenxun
type ResponseAudioTenxun struct {
	Audio string `json:"audio"`
	SessionId string `json:"sessionId"`
	RequestId string `json:"requestId"` // 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId
}

// PublicRequestTenxun
type PublicRequestTenxun struct {
	Action string `json:"action"`
	Region string `json:"region"`
	Timestamp int64 `json:"timestamp"`
	Version   string `json:"version"`
	Authorization string `json:"authorization"`
}

// NewTenxun
func NewTenxun(s *RequestAudioTenxun)(d *RequestAudioTenxun,err error)  {
	if s!=nil{
		d = s
	}else {
		d = new(RequestAudioTenxun)
	}
	if err = d.InitDefault();err!=nil{
		return
	}
	return
}

// 区域
/*
亚太地区(曼谷)	ap-bangkok
华北地区(北京)	ap-beijing
西南地区(成都)	ap-chengdu
西南地区(重庆)	ap-chongqing
华南地区(广州)	ap-guangzhou
华南地区(广州Open)	ap-guangzhou-open
东南亚地区(中国香港)	ap-hongkong
亚太地区(孟买)	ap-mumbai
亚太地区(首尔)	ap-seoul
华东地区(上海)	ap-shanghai
华东地区(上海金融)	ap-shanghai-fsi
华南地区(深圳金融)	ap-shenzhen-fsi
东南亚地区(新加坡)	ap-singapore
欧洲地区(法兰克福)	eu-frankfurt
欧洲地区(莫斯科)	eu-moscow
美国东部(弗吉尼亚)	na-ashburn
美国西部(硅谷)	na-siliconvalley
北美地区(多伦多)	na-toronto
*/

// InitDefault
func (r *RequestAudioTenxun)InitDefault() (err error) {
	if r==nil{
		err = errors.New("requestAudio is nil ")
		return
	}
	if len(r.Action)==0{
		r.Action="TextToVoice"
	}
	if len(r.Version)==0{
		r.Version="2018-05-22"
	}
	if len(r.Region)==0{
		r.Region="ap-guangzhou"
	}
	if r.ModelType==0{
		r.ModelType=1
	}
	if len(r.Codec)==0{
		r.Codec="wav"
	}
	if r.SampleRate==0{
		r.SampleRate = 16000
	}
	if r.PrimaryLanguage==0{
		r.PrimaryLanguage=1
	}
	if len(r.SessionId)==0{
		r.SessionId=uuid.NewV4().String()
	}
	return
}

// PubGetAudioByTextTenxun
func PubGetAudioByTextTenxun(text string, r *RequestAudioTenxun)(ret []byte,err error)  {
	var(
		client = new(http.Client)
		req *http.Request
		now = time.Now().UTC().Format("2006-01-02")
		)
	if req,err = http.NewRequest("GET","",nil);err!=nil{
		return
	}
	req.Header.Set("Authorization",fmt.Sprintf("TC3-HMAC-SHA256 Credential=%s/%s/aai/tc3_request,SignedHeaders=content-type;host, Signature=%s",
		SecretKeyTenxun,now,r.SessionId))
	return
}