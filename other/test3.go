package main

import (
	"fmt"
	"git.yichui.net/tudy/wechat-go/wxweb"
	"regexp"
	"strconv"
	"strings"
)

var (
	_err =fmt.Errorf("%s", "wrong input content")
	err error
	reply string
	sess *wxweb.Session
)
var RegCancelCode = regexp.MustCompile("(取消|撤消)?\\s*([3-9][0-9]{2}[1-7])\\s*(取消|撤消)?")
var RegEmoji = regexp.MustCompile(`<span class="emoji emoji([0-9a-z]+)"></span>`)
func main() {
	s:=pubContentPure("取消 4004")
	fmt.Println(pubParseCancelCode(s))
}



func pubContentPure(in string) (out string) {
	out = RegEmoji.ReplaceAllStringFunc(in, func(iStr string) (oStr string) {
		arr := RegEmoji.FindStringSubmatch(iStr)
		if len(arr) == 2 {
			var outArr []string
			for i := 0; i < len(arr[1]); i++ {
				if (i+1)%5 == 0 {
					_s, _ := strconv.Unquote(`"\U` + "000" + arr[1][i-4:i+1] + `"`)
					outArr = append(outArr, _s)
				}
			}
			oStr = strings.Join(outArr, "")
		}
		return
	})
	return
}

func pubParseCancelCode(con string) (ret string, ret2 string) {
	var (
		code  string
		finds= RegCancelCode.FindAllStringSubmatch(con, -1)
	)
	if len(finds) > 0 {
		if finds[0][1] == "取消" || finds[0][1] == "撤销" {
			ret2 = finds[0][1]
			code = finds[0][2]
		} else if finds[0][len(finds[0])-1] == "取消" || finds[0][len(finds[0])-1] == "撤销" {
			code = finds[0][2]
			ret2 = finds[0][len(finds[0])-1]
		}

	}

	ret = code
	return
}

type LoginForm struct {
	User     string  `form:"user" binding:"required"`
	Password *string `form:"password" binding:"required"`
}

