package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	Str1      = "AccessKeyId=35cbfd79-00e63097-aca04a68-42adb&"
	Str2      = "order-id=1234567890&"
	Str3      = "SignatureMethod=HmacSHA256&"
	Str4      = "SignatureVersion=2&"
	SecretKey = "b1af7228-97b6d63a-1e1044bc-1aa69"
	ApiUrl    = "https://api.huobipro.com/market/history/kline?period=1day&size=1800&symbol=ethusdt&"
)

type klineStruct struct {
	Id     float64 `json:"id"`     //K线id（时间戳）,
	Open   float64 `json:"open"`   //开盘价,
	Close  float64 `json:"close"`  //收盘价,当K线为最晚的一根时，是最新成交价
	Low    float64 `json:"low"`    //最低价,
	High   float64 `json:"high"`   //最高价
	Amount float64 `json:"amount"` //成交量
	Count  float64 `json:"count"`  //成交笔数,
	Vol    float64 `json:"vol"`    //成交额, 即 sum(每一笔成交价 * 该笔的成交量)
}

type KlineArr struct {
	Data []*klineStruct
}

var KlineData *KlineArr

//
func main() {
	var (
		timeNow      = time.Now().UTC()
		client       = &http.Client{}
		request      *http.Request
		response     *http.Response
		responseByte []byte
		timeStamp    string
		secretKeyUrl string
		data         KlineArr
		err          error
	)

	//参数准备
	timeStamp = timeNow.Format("2006-01-02T03:04:05")
	_v := base64.URLEncoding.EncodeToString([]byte(SecretKey))
	secretKeyUrl = url.QueryEscape(_v)
	url := ApiUrl + Str1 + Str2 + Str3 + Str4 + "Timestamp=" + timeStamp + "&" + "Signature=" + secretKeyUrl
	//开始请求
	fmt.Println(url)
	//url1 := "https://api.huobipro.com/market/history/kline?period=1day&size=1800&symbol=ethusdt&AccessKeyId=35cbfd79-00e63097-aca04a68-42adb&order-id=1234567890&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2018-12-21T09:07:19&Signature=YjFhZjcyMjgtOTdiNmQ2M2EtMWUxMDQ0YmMtMWFhNjk%3D"
	//
	if request, err = http.NewRequest("GET", url, nil); err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	if response, err = client.Do(request); err != nil {
		panic(err)
		//go PubGetEthKline()
	}
	if responseByte, err = ioutil.ReadAll(response.Body); err != nil {
		panic(err)
		return
	}
	str := string(responseByte)
	if err = json.Unmarshal([]byte(str), &data); err != nil {
		return
	}
	KlineData = &data
	fmt.Println(KlineData)
}
