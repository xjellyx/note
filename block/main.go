package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/note/block/serve"
)

func main() {
	router := gin.Default()
	router.Handle("POST", "/newBlock", serve.NewBlock)
	router.Handle("GET", "/getBlock/:hash", serve.GetBlock)
	router.Handle("GET", "/getAll", serve.GetAllBlock)
	router.Run("127.0.0.1:8996")
}
