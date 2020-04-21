package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/note/config"
)

var a = &c{
	Name: "dsdasd",
	GG: struct {
		Name     string
		Age      int
		DemoDdqq string
	}{Name: "张三", Age: 9999999999, DemoDdqq: "qqqwewe"},
	StudentAge: 699,
	Data: map[string]interface{}{
		"adada": "8888888888888888",
	},
	Sex:   "boy",
	Class: "012014958",
	Klki:  8888888,
}

func main() {

	var err error
	if err = config.LoadConfigAndSave("全球请求权.yml", a, a); err != nil {
		panic(err)
	}

	var (
		r = gin.Default()
	)

	r.GET("/", getConfig)
	r.Run("0.0.0.0:1563")
	// 挂起进程，直到获取到一个信号
}

func getConfig(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(200, a)
	return
}

type c struct {
	config.Config `yaml:"-"`
	Name          string `yaml:"name" json:"name"`
	GG            struct {
		Name     string
		Age      int
		DemoDdqq string
	} `yaml:"用户"`
	StudentAge int                    `json:"student_age" yaml:"student_age"`
	Data       map[string]interface{} `json:"data" yaml:"data"`
	Sex        string                 `json:"sex" yaml:"sex"`
	Class      string                 `json:"class" yaml:"class"`
	Klki       int
}
