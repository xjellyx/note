package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(copyx("/home/olongfen/win_data/gocode/src/work/ghk-data-back/information-analysis-micro/static/ghk/"+
		"results/c99766c1-aa1b-4848-9cbe-5da7a1ac0082/result/res.jpg", "/home/olongfen/win_data/gocode/src/work/ghk-"+
		"data-back/information-analysis-micro/static/ghk/results/ghk/images/c99766c1-aa1b-4848-9cbe-5da7a1ac0082.jpg"))
}

func copyx(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	println("wwww", err.Error())
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func Zip(srcDir string, zipFileName string) (err error) {
	var (
		zipfile *os.File
	)
	if zipfile, err = os.Create(zipFileName); err != nil {
		return
	}
	defer zipfile.Close()

	// 打开zip文件夹
	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	if err = filepath.Walk(srcDir, func(path string, info fs.FileInfo, err error) error {
		if path == srcDir {
			return nil
		}
		// 获取文件头信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, srcDir+"/")
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		writer, err := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	}); err != nil {
		return
	}
	return

}
