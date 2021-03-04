package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"time"
)

type config struct {
	Dev         bool
	Cookie      string
	Dir         string
	DownloadUrl string
}

var (
	dataMap = make(map[string]map[string]string)
	downMap = make(map[string]string)
	cfg     = new(config)
)

func init() {
	getConfig()
	var (
		d []byte
	)
	d, _ = ioutil.ReadFile("./down_file.json")
	_ = json.Unmarshal(d, &downMap)
}

func main() {

	go func() {
		r := gin.Default()
		r.GET("/file", func(c *gin.Context) {
			c.AbortWithStatusJSON(200, downMap)
		})
		r.GET("/complete", func(c *gin.Context) {
			file := c.Query("filepath")
			if err := os.Remove(cfg.Dir + file); err != nil {
				logrus.Errorln(err)
			}
			c.AbortWithStatusJSON(200, "success")
		})
		r.Run(":8011")
	}()
	var (
		secondParser = cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)

		jobCron = cron.New(cron.WithParser(secondParser), cron.WithChain())
	)
	// 每月1号执行一次
	if _, err := jobCron.AddFunc("0 0 1 * ?", taskDownload); err != nil {
		logrus.Fatalln(err)
	}
	// 每月2号执行一次
	if _, err := jobCron.AddFunc("0 0 2 * ?", mvDownloadFile); err != nil {
		logrus.Fatalln(err)
	}
	jobCron.Start()
	go taskDownload()
	mvDownloadFile()
	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
EXIT:
	for {
		sig := <-sc
		logrus.Infoln("signal: ", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	logrus.Println("exit")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}

func taskDownload() {

	var (
		data []map[string]string
		d    []byte
	)
	if !cfg.Dev {
		d, _ = ioutil.ReadFile("./down_file.json")
		_ = json.Unmarshal(d, &downMap)
		getFileUrl()
		for k, v := range dataMap {
			if _, ok := downMap[k]; ok {
				continue
			}
			data = append(data, v)
		}
	} else {
		data = []map[string]string{{"dir": "shared/2018/Monthly/201806/Tile1_75N180W/", "filename": "SVDNB_npp_20180601-20180630_75N180W_vcmslcfg_v10_c201904251200.tgz", "href": "https://eogdata.mines.edu/wwwdata/viirs_products/dnb_composites/v10//201806/vcmslcfg/SVDNB_npp_20180601-20180630_75N180W_vcmslcfg_v10_c201904251200.tgz", "year": "2018"},
			{"dir": "shared/2018/Monthly/201807/Tile6_00N060E/", "filename": "SVDNB_npp_20180701-20180731_00N060E_vcmslcfg_v10_c201812111300.tgz", "href": "https://eogdata.mines.edu/wwwdata/viirs_products/dnb_composites/v10//201807/vcmslcfg/SVDNB_npp_20180701-20180731_00N060E_vcmslcfg_v10_c201812111300.tgz", "year": "2018"}}

	}

	if len(data) > 0 {
		logrus.Println("开始下载eog文件")
		m := make(map[string]interface{})
		m["data"] = data
		m["download_dir"] = cfg.Dir
		body, _ := json.Marshal(m)
		if req, err := http.NewRequest("POST", cfg.DownloadUrl, bytes.NewReader(body)); err != nil {
			logrus.Errorln(err)
		} else {
			req.Header.Set("Content-Type", "application/json;charset=UTF-8")
			c := http.Client{}
			if resp, err := c.Do(req); err != nil {
				logrus.Errorln(err)
			} else {
				logrus.Infoln(resp.Status)
			}
		}
		logrus.Infoln("异步下载完成")
		for _, v := range data {
			downMap[v["filename"]] = v["dir"]
		}
		saveDownFile()
	} else {
		logrus.Println("没有新文件下载")
	}

}

func getFileUrl() {
	resp, err := http.Get("https://eogdata.mines.edu/pages/download_dnb_composites_iframe.html")
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer resp.Body.Close()
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	d.Find(`#treemenu1>li`).Each(func(i0 int, s0 *goquery.Selection) {
		year := s0.Find("#treemenu1>li>strong").Text()
		monthly := s0.Find("#treemenu1 > li > ul > li:nth-child(2) > strong").Text()
		s0.Find("#treemenu1 > li > ul > li:nth-child(2)>ul>li").Each(func(i1 int, s1 *goquery.Selection) {
			month := s1.Find("#treemenu1 > li > ul > li:nth-child(2)>ul>li>strong").Text()
			s1.Find("#treemenu1 > li > ul > li:nth-child(2)>ul>li>ul>li").Each(func(i2 int, s2 *goquery.Selection) {
				name := s2.Find("#treemenu1 > li > ul > li:nth-child(2)>ul>li>ul>li>strong").Text()
				//fmt.Println(year, monthly, month, name)
				a := s2.Find("a")
				if a.Length() == 2 {
					href, _ := a.Eq(1).Attr("href")
					l := strings.Split(href, "/")
					filename := l[len(l)-1]
					dir := fmt.Sprintf(`/%s/%s/%s/%s/`, year, monthly, month, name)
					dataMap[filename] = map[string]string{}
					dataMap[filename]["href"] = href
					dataMap[filename]["dir"] = dir
					dataMap[filename]["filename"] = filename
					dataMap[filename]["year"] = year
				} else {
					href, _ := a.Attr("href")
					l := strings.Split(href, "/")
					filename := l[len(l)-1]
					dir := fmt.Sprintf(`/%s/%s/%s/%s/`, year, monthly, month, name)
					dataMap[filename] = map[string]string{}
					dataMap[filename]["href"] = href
					dataMap[filename]["dir"] = dir
					dataMap[filename]["filename"] = filename
					dataMap[filename]["year"] = year
				}
			})
		})

	})
}

func mvDownloadFile() {
	var (
		err   error
		files []os.FileInfo
		d     []byte
		m     = make(map[string]string)
	)
	d, _ = ioutil.ReadFile("./down_file.json")
	_ = json.Unmarshal(d, &m)
	if files, err = ioutil.ReadDir(cfg.Dir); err != nil {
		logrus.Errorln(err)
		return
	}
	for _, item := range files {
		if val, ok := m[item.Name()]; ok {
			if !exists(val) {
				if err = os.MkdirAll(cfg.Dir+val, 0777); err != nil {
					logrus.Errorln(err)
					continue
				}
			}
			if err = os.Rename(cfg.Dir+"/"+item.Name(), cfg.Dir+val+item.Name()); err != nil {
				logrus.Errorln(err)
				continue
			}
		}
	}

}

// exists 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

type Reader struct {
	io.Reader
	Total    int64
	Current  int64
	Filename string
}

func (r *Reader) Read(p []byte) (n int, err error) {
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

func getConfig() {

	viper.SetConfigFile("./config.yaml")
	_ = viper.ReadInConfig()
	if err := viper.Unmarshal(cfg); err != nil {
		logrus.Fatalln(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Println("Config file: ", e.Name, " Op: ", e.Op)
		if err := viper.Unmarshal(cfg); err != nil {
			logrus.Fatal(err)
		}
	})

	return
}

func downloadFileProgress(v map[string]string, cookie string) (err error) {
	var (
		client = &http.Client{}
		req    *http.Request
		resp   *http.Response
		fl     os.FileInfo
	)

	if fl, err = os.Stat(cfg.Dir + v["dir"] + v["filename"]); err != nil && !os.IsNotExist(err) {
		logrus.Errorln(err)
		return
	}

	if req, err = http.NewRequest("GET", v["href"], nil); err != nil {
		logrus.Errorln(err)
		return
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Host", "eogdata.mines.edu")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("cache-control", "max-age=0")
	if resp, err = client.Do(req); err != nil {
		logrus.Errorln(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		logrus.Errorln("err: ", resp.Status)
		return
	}
	l := resp.Header.Get("Content-Length")
	ll, _ := strconv.ParseInt(l, 10, 64)
	if ll <= 8888 {
		err = errors.New("got file failed: length >>>> " + l)
		return
	}
	defer resp.Body.Close()
	if fl != nil && fl.Size() >= resp.ContentLength {
		return
	}
	os.MkdirAll(cfg.Dir+v["dir"], 0777)
	f, _err := os.Create(cfg.Dir + v["dir"] + v["filename"])
	if _err != nil {
		err = _err
		return
	}
	defer f.Close()
	r := &Reader{
		Reader:   resp.Body,
		Total:    resp.ContentLength,
		Filename: v["filename"],
	}
	io.Copy(f, r)
	downMap[v["filename"]] = v["dir"]
	saveDownFile()
	//
	return
}
func saveDownFile() {
	d, _ := json.Marshal(downMap)
	f, _ := os.Create("./down_file.json")
	f.Write(d)
	logrus.Println("保存已下载文件信息完成")
}
