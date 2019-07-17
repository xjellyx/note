package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type KlineStruct struct {
	Id     string `json:"id"`     //K线id（时间戳）,
	Open   string `json:"open"`   //开盘价,
	Close  string `json:"close"`  //收盘价,当K线为最晚的一根时，是最新成交价
	Low    string `json:"low"`    //最低价,
	High   string `json:"high"`   //最高价
	Amount string `json:"amount"` //成交量
	Count  string `json:"count"`  //成交笔数,
	Vol    string `json:"vol"`    //成交额, 即 sum(每一笔成交价 * 该笔的成交量)
}

const (
	str1 = "SignatureMethod=HmacSHA256&"
	str2 = "SignatureVersion=2&"
	str3 = "order-id=1234567890&"
	str4 = "AccessKeyId=35cbfd79-00e63097-aca04a68-42adb&"
)

func getBody() (ret []byte, err error) {
	var (
		apiUrl    string
		timeNow   = time.Now().UTC()
		secretKet = "b1af7228-97b6d63a-1e1044bc-1aa69"
		client    = &http.Client{}
		reqest    *http.Request
		resp      *http.Response
		data      []byte
	)
	Timestamp0 := timeNow.Format("2006-01-02T03:04:05")
	//fmt.Printf("%s", Timestamp0)
	apiUrl = "https://api.huobipro.com/market/history/kline?period=15min&size=1800&symbol=ethusdt&"
	uEnc := base64.URLEncoding.EncodeToString([]byte(secretKet))
	keyUrl := url.QueryEscape(uEnc)
	url := apiUrl + str4 + str3 + str1 + str2 + "Timestamp=" + Timestamp0 + "&" + "Signature=" + keyUrl

	reqest, err = http.NewRequest("GET", url, nil)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML,"+
		" like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	if err != nil {
		panic(err)
		return
	}
	reqest.Header.Add("authority", "api.huobipro.com")
	reqest.Header.Add("upgrade-insecure-requests", strconv.Itoa(1))
	reqest.Header.Add("scheme", "https")
	reqest.Header.Add("accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b`)
	if resp, err = client.Do(reqest); err != nil {
		panic(err)
		return
	}
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
		return
	}
	ret = data
	return
}

func main() {
	var dict = map[string]interface{}{}
	s, _ := getBody()
	str := string(s)

	//ss1 := strings.TrimRight(ss, "}")
	dict[str] = `"status":"ok"`
	fmt.Println(dict)

}
