package serve

import (
	"bytes"
	"encoding/json"
	"git.yichui.net/tudy/wechat-go/serve/setting"
	"github.com/suboat/sorm/log"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	Bot1 = UserInfo{
		ApiKey: "68c694ab953a4a2ca29eb23db90840f1",
		UserId: "5279aef0522a90de",
	}
	Bot2 = UserInfo{
		ApiKey: "b482bea7b7944358a61d0dfe393a76bd",
		UserId: "b7d9c64bb9361593",
	}
	Bot3 = UserInfo{
		ApiKey: "2c4ddd3c22864661b2b7abc716068793",
		UserId: "1ff38ff7dd42d924",
	}
	Bot4 = UserInfo{
		ApiKey: "d6d85fffd0974bc092dc634ac9c24775",
		UserId: "dee2f0f2a30c70fd",
	}
	Bot5 = UserInfo{
		ApiKey: "6915d16fbece46c7a1960952cbbcc9c6",
		UserId: "4ee9d81bc165928f",
	}
	Bot6 = UserInfo{
		ApiKey: "172b62aa346b4bbfa12fc0f13d9e49cc",
		UserId: "6abf66ce7e198ec6",
	}
	Bot7 = UserInfo{
		ApiKey: "bb9e8e848c874f8ca040e9e45a2212c6",
		UserId: "734d7c964ca2b364",
	}
	//
	Bots       = []UserInfo{Bot1, Bot2, Bot3, Bot4, Bot5, Bot6, Bot7}
	index      = 0
	ErrorReply = ""
	count      = 0
)

// GetBotReply 获取机器人的回答
func GetBotReply(content string) (ret string, err error) {
	var (
		url     = setting.Settings.ChatBot.PostUrl
		client  = &http.Client{}
		dataReq = &BotRequest{
			ReqType: 0,
			Perception: Perception{
				InputText: InputText{
					Text: content,
				},
			},
		}
		dataResp struct {
			Results []interface{}
			Intent  struct {
				Code int
			}
		}
		r string
	)

	// 设置请求参数
	botNow := Bots[index]
	dataReq.UserInfo = botNow
	bs, _ := json.Marshal(dataReq)
	body := bytes.NewBuffer(bs)
	req, _ := http.NewRequest("POST", url, body)
	// 发送请求
	if resp, _err := client.Do(req); _err != nil {
		err = _err
		log.Errorf(`client Do error:%v`, err)
		return
	} else {
		// 读取返回的body
		bss, _ := ioutil.ReadAll(resp.Body)
		// 关闭body
		resp.Body.Close()
		// 解析成json
		json.Unmarshal(bss, &dataResp)
	}
	// 解析interface
	for _, v := range dataResp.Results {
		if _name, _ok := v.(map[string]interface{}); _ok {
			if name, ok := _name["values"]; ok {
				if name1, ok1 := name.(map[string]interface{}); ok1 {
					if _ret, ok2 := name1["text"]; ok2 {
						r = _ret.(string)
					}
				}
			}
		}
	}
	//
	count++

	// 如果发送失败则重新尝试
	if dataResp.Intent.Code != 10004 {
		index++
		if index >= len(Bots) {
			index = 0
		}
		GetBotReply(content)
		// 休息一下，避免被ban
		time.Sleep(time.Second)
	}

	if dataResp.Intent.Code != 10004 {
		// 故障时回复
		r = ErrorReply
	}
	ret = r
	return
}

type BotRequest struct {
	ReqType    int        `json:"reqType"`
	Perception Perception `json:"perception"`
	UserInfo   UserInfo   `json:"userInfo"`
}

type UserInfo struct {
	ApiKey string `json:"apiKey"`
	UserId string `json:"userId"`
}

type Perception struct {
	InputText InputText `json:"inputText"`
}

type InputText struct {
	Text string `json:"text"`
}
