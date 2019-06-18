package ctrl

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"time"
)

func InitRouter(app *iris.Application) (err error) {
	s := sessions.New(sessions.Config{
		Cookie:  "sessionCookie",
		Expires: 24 * time.Hour,
	})

	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		s.Start,
	)
	admin.Handle(AdminCTRL{})
	return
}
