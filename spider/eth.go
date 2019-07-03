package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	var (
		res    *http.Response
		reqest *http.Request
		client = &http.Client{}
		err    error
	)
	reqest, err = http.NewRequest("POST", "http://srh.bankofchina.com/search/whpj/search.jsp", strings.NewReader("pjname=1316"))
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like "+
		"Gecko) Chrome/73.0.3683.75 Safari/537.36")
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Add("Referer", "http://www.boc.cn/sourcedb/whpj/")
	reqest.Header.Add("Origin", "http://srh.bankofchina.com")
	reqest.Header.Add("Cookie", "JSESSIONID=0000poVLE_MQLZrCgwhZtvexMYX:-1")
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,"+
		"image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	reqest.Header.Add("Host", "srh.bankofchina.com")
	if res, err = client.Do(reqest); err != nil {
		return
	}
	defer res.Body.Close()
	test()
}

func test() {
	var (
		doc *goquery.Document
	)
	resp, err := http.Get("https://www.feixiaohao.com/currencies/ethereum/")

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	if doc, err = goquery.NewDocumentFromReader(resp.Body); err != nil {
		return
	}
	//ll := doc.Find(`.priceInfo .convert`).Text()
	//fmt.Println(ll)
	doc.Find(`.priceInfo`).Each(func(i int, selection *goquery.Selection) {
		t := selection.Find(".convert").Text()
		fmt.Println(t)
		a := ""
		for _i, v := range t {
			if string(v) == "$" {
				a = t[:_i]
			}
		}
		fmt.Println(a)

	})
}
func priGetEthPrice() (ret float64, err error) {
	var (
		url  = "http://api.zb.cn/data/v1/ticker?market=eth_usdt"
		resp *http.Response
		doc  *goquery.Document
		data float64
	)
	// 获取数据
	if resp, err = http.Get(url); err == nil {
		defer resp.Body.Close()
	} else {
		return
	}
	// 解析参数
	if doc, err = goquery.NewDocumentFromReader(resp.Body); err != nil {
		return
	}
	var d *struct {
		Ticker struct {
			Buy string `json:"buy"`
		}
	}

	if err = json.Unmarshal([]byte(doc.Text()), &d); err != nil {
		return
	}
	if data, err = strconv.ParseFloat(d.Ticker.Buy, 64); err != nil {
		return
	}

	ret = data

	return
}
