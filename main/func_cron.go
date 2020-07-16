package main

import "github.com/gin-gonic/gin"

type TaskTime struct {
	Time string `json:"time" form:"time" binding:"required"`
	Week int    `json:"week" form:"week" binding:"required"`
}

func main()  {
d:=gin.Default()
d.POST("/", func(c *gin.Context) {
	var a struct{
		Email string `json:"email" binding:"email" validate:"required,email"`
		Run TaskTime `json:"run" form:"run" binding:"required"`
	}
	if err := c.ShouldBind(&a);err!=nil{
		c.AbortWithError(404,err)
		return
	}
	c.JSON(200,gin.H{"msg":"success"})
})
	d.GET("/", func(c *gin.Context) {
		var a struct{
			Email string `json:"email" binding:"email" validate:"required,email"`
		}
		if err := c.ShouldBind(&a);err!=nil{
			c.AbortWithError(404,err)
			return
		}
		c.JSON(200,gin.H{"msg":"success"})
	})

d.Run(":9090")

}

