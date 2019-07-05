package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string  `form:"user" binding:"required"`
	Password *string `form:"password" binding:"required"`
}

func main() {
	fmt.Println(3 << 3)
}

func startPage(c *gin.Context) {
	var person LoginForm
	if c.ShouldBind(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.User)
		log.Println(person.Password)
		c.JSON(200, person)
	} else {
		c.String(400, c.ShouldBindJSON(&person).Error())
	}

}
