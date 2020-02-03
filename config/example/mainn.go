package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olefen/note/config"
)

var a = &c{
	Name: "dsdasd",
	GG: struct {
		Name string
		Age  int
	}{Name: "张三", Age: 19},
	StudentAge: 699,
}

func main() {

	var err error
	if err = config.LoadConfiguration("全球请求权.yml", a, a); err != nil {
		panic(err)
	}
	a.SetHookChange(a.Save)
	if err = a.Save(nil); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/", getConfig)
	r.Run("0.0.0.0:1563")

}

func getConfig(ctx *gin.Context) {
	if err := config.LoadConfiguration("全球请求权.yml", a, a); err != nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"msg": err,
		})
		return
	}
	ctx.AbortWithStatusJSON(200, a)
	return
}

type c struct {
	config.Config `yaml:"-"`
	Name          string `yaml:"name" json:"name"`
	GG            struct {
		Name string
		Age  int
	} `yaml:"用户"`
	StudentAge int                    `json:"student_age" yaml:"student_age"`
	Data       map[string]interface{} `json:"data" yaml:"data"`
}
