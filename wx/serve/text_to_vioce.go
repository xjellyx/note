package serve

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	aai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai/v20180522"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"io/ioutil"
	"net/http"
)

var (
	ApiKeyBaidu     = "ZHRLdu2E8FgqvNdPVh4e8H99"
	SecretKeyBaidu  = "ESceg2zGaGi5dEr07kZVmnl2Ge7OXKg3"
	TokenApi        = `https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&`
	AudioApiBaidu   = "https://tsn.baidu.com/text2audio"
	SecretIdTenxun  = "AKIDV7DNUB5lZDeB7ru4mvSfzEibJ2JRszky"
	SecretKeyTenxun = "Ly1jovx4gYyAAlndqap4nYQHA12o3dcO"
)

// PubGetToken 获取token
func PubGetToken(apiKey, secretKey string) (ret ResponseTokenBaidu, err error) {
	var (
		data ResponseTokenBaidu
		body []byte
		resp *http.Response
	)
	api := fmt.Sprintf(TokenApi+`client_id=%s&client_secret=%s`, apiKey, secretKey)
	// 获取token
	if resp, err = http.Get(api); err != nil {
		return
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return
	}
	ret = data
	return
}

// ResponseTokenBaidu 获取百度token
type ResponseTokenBaidu struct {
	AccessToken   string `json:"access_token"`
	ExpiresIn     int64  `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}

// RequestAudio 请求参数
type RequestAudioBaidu struct {
	Tex  string `json:"tex"`  // 文本
	Tok  string `json:"tok"`  // token
	CUID string `json:"cuid"` // 用户唯一标识
	Lan  string `json:"lan"`  // zh
	Ctp  int    `json:"ctp"`  // 客户端类型选择，web端填写固定值1
	Spd  int    `json:"spd"`  // 语速，取值0-15，默认为5中语速
	Pit  int    `json:"pit"`  // 音调，取值0-15，默认为5中语调
	Vol  int    `json:"vol"`  // 音量，取值0-15，默认为5中音量
	Per  int    `json:"per"`  // 发音人选择, 0为普通女声，1为普通男生，3为情感合成-度逍遥，4为情感合成-度丫丫，默认为普通女声
	Aue  int    `json:"aue"`  // 3为mp3格式(默认)； 4为pcm-16k；5为pcm-8k；6为wav（内容同pcm-16k）; 注意aue=4或者6是语音识别要求的格式， // 但是音频内容不是语音识别要求的自然人发音，所以识别效果会受影响。
}

// NewAudioBaidu
func NewAutioBaidu(s *RequestAudioBaidu) (d *RequestAudioBaidu, err error) {
	if s != nil {
		d = s
	} else {
		d = new(RequestAudioBaidu)
	}
	if err = d.InitDefault(); err != nil {
		return
	}
	return
}

// InitDefault
func (r *RequestAudioBaidu) InitDefault() (err error) {
	if r == nil {
		err = errors.New("requestAudio is nil ")
		return
	}
	if len(r.CUID) == 0 {
		r.CUID = uuid.NewV4().String()
	}
	if len(r.Lan) == 0 {
		r.Lan = "zh"
	}
	if r.Ctp == 0 {
		r.Ctp = 1
	}
	if r.Spd == 0 {
		r.Spd = 5
	}
	if r.Pit == 0 {
		r.Pit = 5
	}
	if r.Vol == 0 {
		r.Vol = 5
	}
	if r.Aue == 0 {
		r.Aue = 6
	}
	return
}

// PubPostAudioByTextBaidu 获取语音
func PubPostAudioByTextBaidu(text string, t *ResponseTokenBaidu) (ret []byte, err error) {
	var (
		client = &http.Client{}
		resp   *http.Response
		rb     *RequestAudioBaidu
	)

	if rb, err = NewAutioBaidu(nil); err != nil {
		return
	}
	rb.Tex = text
	rb.CUID = uuid.NewV4().String()
	rb.Tok = t.AccessToken
	r, _ := json.Marshal(rb)
	req, _err := http.NewRequest("POST", AudioApiBaidu, bytes.NewReader(r))
	if _err != nil {
		err = _err
		return
	}
	if resp, err = client.Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	if ret, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	return
}

// PubGetAudioByTextBaidu 获取语音
func PubGetAudioByTextBaidu(text string, t *ResponseTokenBaidu) (ret []byte, err error) {
	var (
		resp *http.Response
		rb   *RequestAudioBaidu
	)

	if rb, err = NewAutioBaidu(nil); err != nil {
		return
	}
	rb.CUID = uuid.NewV4().String()
	rb.Tok = t.AccessToken

	if resp, err = http.Get(fmt.Sprintf(`%s?lan=zh&ctp=1&cuid=%s&tok=%s&tex=%s`, AudioApiBaidu, rb.CUID,
		rb.Tok, text)); err != nil {
		return
	}

	defer resp.Body.Close()

	if ret, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	return
}

// RequestAudioTenxun 请求参数
type RequestAudioTenxun struct {
	// 必填
	Text      string `json:"text"`
	SecretId  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
	SessionId string `json:"sessionId"`
	ModelType int    `json:"modelType"`
	// 选填
	Volume          *float32 `json:"volume"`          // 音量大小[0-10],默认0
	Speed           *float32 `json:"speed"`           // 语速, [-2,2],2代表0.6倍 -1代表0.8倍 0代表1.0倍（默认） 1代表1.2倍 2代表1.5倍
	ProjectId       *int     `json:"projectId"`       // 项目id,默认0
	VoiceType       *int     `json:"voiceType"`       // 音色 0-亲和女声(默认) 1-亲和男声 2-成熟男声 3-活力男声 4-温暖女声 5-情感女声 6-情感男声
	PrimaryLanguage *int     `json:"primaryLanguage"` // 语言类型, 1 中文,默认;2-英文
	SampleRate      *int     `json:"sampleRate"`      // 音频采样率 默认16k, 8000: 8k
	Codec           *string  `json:"codec"`           // wav默认; mp3

}

// Valid
func (r *RequestAudioTenxun) Valid() (err error) {
	if r == nil {
		err = errors.New("RequestAudioTenxun is nil")
		return
	}
	if len(r.Text) == 0 {
		err = errors.New("wrong text")
		return
	}
	if len(r.SecretId) == 0 {
		err = errors.New("wrong secretId")
		return
	}
	if len(r.SecretKey) == 0 {
		err = errors.New("wrong secretKey")
		return
	}
	if len(r.Region) == 0 {
		err = errors.New("region is nil")
		return
	}
	if r.ModelType == 0 {
		r.ModelType = 1
	}
	if len(r.SessionId) == 0 {
		r.SessionId = uuid.NewV4().String()
	}

	return
}

// ToMapDATA
func (r *RequestAudioTenxun) ToMapDATA() (ret map[string]interface{}, err error) {

	if err = r.Valid(); err != nil {
		return
	}

	var (
		paramMap = make(map[string]interface{})
	)

	paramMap["text"] = r.Text
	paramMap["SessionId"] = r.SessionId
	paramMap["ModelType"] = r.ModelType
	if r.Volume != nil && *r.Volume > 0 {
		paramMap["volume"] = r.Volume
	}
	if r.Volume != nil && *r.Speed > 0 {
		paramMap["speed"] = r.Speed
	}
	if r.ProjectId != nil && *r.ProjectId > 0 {
		paramMap["projectId"] = r.ProjectId
	}
	if r.VoiceType != nil && *r.VoiceType > 0 {
		paramMap["voiceType"] = r.VoiceType
	}
	if r.PrimaryLanguage != nil && *r.PrimaryLanguage > 0 {
		paramMap["primaryLanguage"] = r.PrimaryLanguage
	} else {
		paramMap["primaryLanguage"] = PrimaryLanguage
	}
	if r.SampleRate != nil && *r.SampleRate > 0 {
		paramMap["sampleRate"] = r.SampleRate
	} else {
		paramMap["sampleRate"] = SampleRate
	}
	if r.Codec != nil && len(*r.Codec) > 0 {
		paramMap["codec"] = r.Codec
	} else {
		paramMap["codec"] = CodecWav
	}

	ret = paramMap
	return
}

// PubGetAudioByTextTenxun 获取语音
func PubGetAudioByTextTenxun(r *RequestAudioTenxun) (ret []byte, err error) {

	if err = r.Valid(); err != nil {
		return
	}

	var (
		client   *aai.Client
		paramMap = make(map[string]interface{})
		resp     *aai.TextToVoiceResponse
	)
	if paramMap, err = r.ToMapDATA(); err != nil {
		return
	}
	// 创建客户端请求
	if client, err = aai.NewClient(common.NewCredential(r.SecretId, r.SecretKey), r.Region,
		profile.NewClientProfile()); err != nil {
		return
	}

	req := aai.NewTextToVoiceRequest()
	params, _ := json.Marshal(paramMap)
	if err = req.FromJsonString(string(params)); err != nil {
		return
	}

	// 获取参数
	if resp, err = client.TextToVoice(req); err != nil {
		return
	}

	decodeBytes, _ := base64.StdEncoding.DecodeString(*resp.Response.Audio)

	ret = decodeBytes
	return
}
