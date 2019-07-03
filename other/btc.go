package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"path/filepath"
)

func main() {
	var (
		body   = &bytes.Buffer{}
		writer = multipart.NewWriter(body)
		img    = bytes.NewBuffer([]byte("http://img0.imgtn.bdimg.com/it/u=3440048196,1495127308&fm=26&gp=0.jpg"))
		fw     io.Writer
		//resp   *http.Response
		err error
	)

	// 写入第1部分信息
	if fw, err = writer.CreateFormFile("file", filepath.Base("wx_icon.jpg")); err != nil {
		return
	}
	// 写入第2部分信息
	if _, err = io.Copy(fw, img); err != nil {
		return
	}
	// 写入第3部分信息
	if err = writer.Close(); err != nil {
		return
	}

}
