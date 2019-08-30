package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.yichui.net/tudy/go-rest"
	"github.com/tealeg/xlsx"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"
)

type d struct {
	a string
	b string
	c string
	d string
}

func main() {
	var (
		file  *xlsx.File
		sheet *xlsx.Sheet
		row   *xlsx.Row
		cell  *xlsx.Cell
		err   error
		datas = []d{}
	)
	datas = append(datas, d{"a1", "b1", "c1", "d1"})
	datas = append(datas, d{"a2", "b2", "c2", "d2"})
	datas = append(datas, d{"a13", "b13", "c13", "d13"})
	file = xlsx.NewFile()
	if sheet, err = file.AddSheet("Sheet1"); err != nil {
		return
	}
	row = sheet.AddRow()
	// 第一列
	cell = row.AddCell()
	cell.Value = "标准问题"
	// 第二列
	cell = row.AddCell()
	cell.Value = "分类"
	// 第三列
	cell = row.AddCell()
	cell.Value = "答案"
	// 第四列
	cell = row.AddCell()
	cell.Value = "相似问题"
	for _, v := range datas {
		_row := sheet.AddRow()
		cell = _row.AddCell()
		cell.Value = v.a

		cell = _row.AddCell()
		cell.Value = v.b

		cell = _row.AddCell()
		cell.Value = v.c

		cell = _row.AddCell()
		cell.Value = v.d
	}

	if err = file.Save("demo.xlsx"); err != nil {
		return
	}
	fileDir := "/home/srlemon/金融.xlsx"
	b, _err := ioutil.ReadFile(fileDir)
	if _err != nil {
		panic(_err)
	}

	s, _ := PubHashFromFilepool(b)
	d := rest.PubJsonMust(s)
	var _d struct {
		Data struct {
			Hash string
		}
	}
	if err = json.Unmarshal([]byte(d), &_d); err != nil {
		return
	}

	fmt.Println(_d.Data.Hash)

}

// PubHashFromFilepool 上传文件获取hash
func PubHashFromFilepool(input []byte) (ret string, err error) {
	start := time.Now()
	var (
		body      = &bytes.Buffer{}
		inputBody = bytes.NewBuffer(input)
		writer    = multipart.NewWriter(body)
		resp      *http.Response
		fw        io.Writer
		url       = "https://api.yichui.net/api/yichui/filepool/upload-file?form=file"
		data      struct {
			Data struct {
				Hash string
			}
		}
	)

	// 写入FormFile信息
	if fw, err = writer.CreateFormFile("file", filepath.Base(url)); err != nil {
		return
	}

	fmt.Println(filepath.Base(url))

	// 拷贝内容
	if _, err = io.Copy(fw, inputBody); err != nil {
		return
	}

	// 关闭writer
	if err = writer.Close(); err != nil {
		return
	}

	if resp, err = http.Post(url, writer.FormDataContentType(), body); err != nil {
		return
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(bs))
	if err = json.Unmarshal(bs, &data); err != nil {
		return
	}
	ret = data.Data.Hash
	end := time.Now()
	fmt.Println("filepool return costs:", end.Sub(start).Seconds())
	return
}
