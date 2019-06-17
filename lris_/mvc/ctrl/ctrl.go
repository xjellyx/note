package ctrl

import (
	"github.com/LnFen/note/lris-learn/mvc"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type AdminCTRL struct {
	Ctx     iris.Context
	Session sessions.Session
}

func (a *AdminCTRL) PostLogin(ctx iris.Context) (ret mvc.Result, err error) {
	var form *project.FormLogin
	if err = a.Ctx.ReadJSON(form); err != nil {
		return
	}
	if form.Name == "" || form.Password == "" {
		ret = mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": false,
				"message": "用户名或密码为空,请重新填写后尝试登录",
			},
		}
		return
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "1",
				"success": "登录成功",
				"message": "管理员登录成功",
			},
		}, nil
	}
	return
}
