package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	engine := gin.Default()
	engine.POST("/", func(ctx *gin.Context) {
		fh, _ := ctx.FormFile("file")
		f, _ := fh.Open()
		fs, _ := os.OpenFile("demo.tif", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm)
		defer f.Close()
		defer fs.Close()
		reader := bufio.NewReader(f)
		writer := bufio.NewWriter(fs)
		buf := make([]byte, 1024*1024*3) // 3M buf
		for {
			n, err := reader.Peek(0)
			if err == io.EOF {
				writer.Flush()
				break
			} else if err != nil {
				return
			} else {
				writer.Write(buf[:n])
			}
		}
	})
}
