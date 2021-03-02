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
	"strings"
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
	t := time.NewTicker(time.Second * 60 * 10)
	for {
		select {
		case <-t.C:
			taskDown()
		}
	}
	//	var state int32 = 1
	//	sc := make(chan os.Signal, 1)
	//	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//EXIT:
	//	for {
	//		sig := <-sc
	//		logrus.Infoln("signal: ", sig.String())
	//		switch sig {
	//		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	//			atomic.StoreInt32(&state, 0)
	//			break EXIT
	//		case syscall.SIGHUP:
	//		default:
	//			break EXIT
	//		}
	//	}
	//
	//	logrus.Println("exit")
	//	time.Sleep(time.Second)
	//	os.Exit(int(atomic.LoadInt32(&state)))
}

func taskDown() {
	var (
		resp *http.Response
		err  error
		body []byte
		m    = make(map[string]map[string]string)
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

	for _, v := range m {
		go downloadFile(v)
	}
}

func downloadFile(v map[string]string) {
	var (
		err    error
		client = &http.Client{}
		req    *http.Request
		resp   *http.Response
		fl     os.FileInfo
	)
	list := strings.Split(v["dir"], "shared")
	if fl, err = os.Stat("." + list[1] + v["filename"]); err != nil && !os.IsNotExist(err) {
		logrus.Errorln(err)
		return
	}

	if req, err = http.NewRequest("GET", "http://bw.imeizi.ml:8000/financial.starwiz.cn"+list[1]+v["filename"], nil); err != nil {
		logrus.Errorln(err)
		return
	}
	if resp, err = client.Do(req); err != nil {
		logrus.Errorln(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		logrus.Errorln("err: ", resp.Status)
		return
	}
	defer resp.Body.Close()
	if fl != nil && fl.Size() == resp.ContentLength {
		return
	}
	os.MkdirAll("."+list[1], 0777)
	f, _err := os.Create("." + list[1] + v["filename"])
	if _err != nil {
		logrus.Errorln(_err)
		return
	}
	defer f.Close()
	r := &Reader1{
		Reader:   resp.Body,
		Total:    resp.ContentLength,
		Filename: v["filename"],
	}
	io.Copy(f, r)

	http.Get("http://45.79.100.123:8011/complete?filepath=" + v["dir"] + v["filename"])
	return
}
