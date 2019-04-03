package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type EthData struct {
	//eth数据结构体
	Id                string `json:"id"`
	Name              string `json:"name"`           // 名字
	Symbol            string `json:"symbol"`         // 标记符号
	PriceUsd          string `json:"price_usd"`      // 美元价格
	PriceBtc          string `json:"price_btc"`      // 比特币比例
	VolumeUsd         string `json:"24h_volume_usd"` //24小时成交量
	MarketCapUsd      string `json:"market_cap_usd"`
	AvailableSupply   string `json:"available_supply"`   // 可得到供应
	TotalSupply       string `json:"total_supply"`       //总供应
	PercentChange_1h  string `json:"percent_change_1h"`  //近一小时价格变动
	PercentChange_24h string `json:"percent_change_24h"` //
	PercentChange_7d  string `json:"percent_change_7d"`  //
	LastUpdated       string `json:"last_updated"`       //最近更新
}

func getUrl() (ret string, err error) {
	var (
		url  string
		resp *http.Response
		data []byte
	)
	url = "https://api.coinmarketcap.com/v1/ticker/ethereum/"
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	ret = string(data)
	return
}
func main() {
	a, _ := getUrl()
	fmt.Println(a)
	var ss EthData
	b := strings.Trim(a, "[]")
	json.Unmarshal([]byte(b), &ss)
	fmt.Println(ss)
	fmt.Printf("%T\n", ss)
	d := "0.026"
	fmt.Println(strconv.ParseFloat(d, 64))

}
