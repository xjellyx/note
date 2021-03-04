package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Reader1 struct {
	io.Reader
	Total    int64
	Current  int64
	Filename string
}

func (r *Reader1) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	fmt.Printf("下载 %s 进度  %.2f%% \n", r.Filename, float64(r.Current*10000/r.Total)/100)
	var (
		complete = float64(100)
	)
	if float64(r.Current*10000/r.Total)/100 == complete {
		//
	}

	return
}

func main() {
	taskDown()
	t := time.NewTicker(time.Second * 60 * 60)
	for {
		select {
		case <-t.C:
			taskDown()
		}
	}
}

func taskDown() {
	var (
		resp *http.Response
		err  error
		body []byte
		m    = make(map[string]string)
	)
	if resp, err = http.Get("http://45.79.100.123:8011/file"); err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(body, &m); err != nil {
		log.Fatalln(err)
	}
	for k, v := range m {
		go downloadFile(k, v)
	}
}

func downloadFile(filename, dir string) {
	var (
		err      error
		client   = &http.Client{}
		req      *http.Request
		resp     *http.Response
		fl       os.FileInfo
		filepath = "~/data/eogdata" + dir + filename
	)
	if fl, err = os.Stat(filepath); err != nil && !os.IsNotExist(err) {
		logrus.Errorln(err)
		return
	}

	if req, err = http.NewRequest("GET", "http://bw.imeizi.ml:8000/financial.starwiz.cn"+dir+filename, nil); err != nil {
		logrus.Errorln(err)
		return
	}
	if resp, err = client.Do(req); err != nil {
		logrus.Errorln(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		logrus.Errorln("err: ", resp.Status, "\n", dir+filename)
		return
	}
	defer resp.Body.Close()
	if fl != nil && fl.Size() == resp.ContentLength {
		return
	}
	os.MkdirAll("~/data/eogdata"+dir, 0777)
	f, _err := os.Create(filepath)
	if _err != nil {
		logrus.Errorln(_err)
		return
	}
	defer f.Close()
	r := &Reader1{
		Reader:   resp.Body,
		Total:    resp.ContentLength,
		Filename: filename,
	}
	io.Copy(f, r)

	http.Get("http://45.79.100.123:8011/complete?filepath=" + dir + filename)
	return
}
