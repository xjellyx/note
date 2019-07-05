package serve

import (
	"encoding/json"
	"git.yichui.net/tudy/wechat-go/wxweb"
	"github.com/suboat/sorm/log"
	"regexp"
	"strings"
	"sync"
)

var lockInitGroup = sync.RWMutex{}
var (
	CacheGroupToWxName = make(map[string]string)
	RegWxName          = regexp.MustCompile("^[0-9]{6}$")
)

// Register
func Register(sess *wxweb.Session) (err error) {
	if err = sess.HandlerRegister.Add(wxweb.MSG_INIT, handlerInit, "Init"); err != nil {
		return
	}
	if err = sess.HandlerRegister.Add(wxweb.MSG_TEXT, handlerTextFriend, "InitFriend"); err != nil {
		return
	}

	if err = sess.HandlerRegister.EnableByName("Init"); err != nil {
		return
	}
	if err = sess.HandlerRegister.EnableByName("InitFriend"); err != nil {
		return
	}
	return
}

// handlerInit
func handlerInit(sess *wxweb.Session, data *wxweb.ReceivedMessage) {
	var (
		recentArr []string
		groupArr  []string
		friendArr []string
		err       error
	)

	if v, ok := data.Raw["StatusNotifyUserName"]; ok {
		switch t := v.(type) {
		case string:
			recentArr = strings.Split(t, ",")
		}
	}

	if len(recentArr) > 0 {
		for _, v := range recentArr {
			if strings.Contains(v, "@@") {
				groupArr = append(groupArr, v)
			} else if len(v) > 0 {
				if string(v[0]) == "@" && string(v[1]) != "@" {
					friendArr = append(friendArr, v)
				}
			}
		}
	}
	if err = initFriend(sess, friendArr, false); err != nil {
		log.Error(err)
	}

}

// initFriend
func initFriend(sess *wxweb.Session, friendArr []string, isHandler bool) (err error) {
	lockInitGroup.Lock()
	defer lockInitGroup.Unlock()
	var (
		loopBatch = 10 // 每次拉取10个
		friendMap = make(map[string]int)
		paramArr  []*wxweb.User
	)
	val := sess.Cm.GetAll()
	if len(val) > 0 {
		for _, v := range sess.Cm.GetAll() {
			if string(v.UserName[0]) == "@" && string(v.UserName[0]) != "@" {
				friendMap[v.UserName] += 1
			}
		}
	}
	if len(friendArr) > 0 {
		for _, v := range friendArr {
			if strings.Contains(v, "@") {
				friendMap[v] += 1
			}
		}
	}

	// 整理请求参数
	for k, _ := range friendMap {
		if _, ok := CacheGroupToWxName[k]; ok == false || isHandler == true {
			paramArr = append(paramArr, &wxweb.User{UserName: k})
		}
	}

	for len(paramArr) > 0 {
		//
		var params []*wxweb.User
		if len(paramArr) >= loopBatch {
			params = paramArr[0:loopBatch]
			paramArr = paramArr[loopBatch:]
		} else {
			params = paramArr
			paramArr = nil
		}
		//
		if d, _err := sess.Api.WebWxBatchGetContact(sess.WxWebCommon, sess.WxWebXcg, sess.GetCookies(), params); _err != nil {
			err = _err
			log.Error(err)
			return
		} else {
			var (
				resp = new(wxweb.WxWebBatchGetContactResponse)
			)
			if err = json.Unmarshal(d, resp); err != nil {
				return
			}
			for _, v := range resp.ContactList {
				if len(v.MemberList) == 0 {
					log.Warnf(`[group-empty-member] %s -> %s`, v.UserName, v.NickName)
					continue
				}

			}
		}
	}

	return
}

// handlerTextFriend
func handlerTextFriend(sess *wxweb.Session, msg *wxweb.ReceivedMessage) {
	var (
		err   error
		reply string
	)
	if msg.FromUserName == sess.Bot.UserName {
		log.Println(msg.FromUserName, "============", sess.Bot.UserName)
		return
	}
	if msg.IsGroup {
		log.Printf(`"%s" 是一个群`, msg.GroupName)
		return
	}
	if msg.RecommendInfo.NickName == `A罩杯` {
		log.Println("这个人很毒辣,别招惹她")
		return
	}

	if reply, err = GetBotReply(msg.Content); err != nil {
		log.Error("[GetBotReply] err: ", err)
	}

	if len(reply) > 0 {
		if _, _, err = sess.SendText(reply, sess.Bot.UserName, msg.FromUserName); err != nil {
			log.Error("[sedText] err: ", err)
		}
	}

}
