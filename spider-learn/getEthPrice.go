package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	type ethData struct {
		//eth数据结构体
		CreateTime        time.Time `sorm:"index" json:"createTime,omitempty"` // 创建时间
		UpdateTime        time.Time `json:"update_time"`
		Id                string    `json:"id"`
		Name              string    `json:"name"`           // 名字
		Symbol            string    `json:"symbol"`         // 标记符号
		PriceUsd          string    `json:"price_usd"`      // 美元价格
		PriceBtc          string    `json:"price_btc"`      // 比特币比例
		VolumeUsd         string    `json:"24h_volume_usd"` //24小时成交量
		MarketCapUsd      string    `json:"market_cap_usd"`
		AvailableSupply   string    `json:"available_supply"`   // 可得到供应
		TotalSupply       string    `json:"total_supply"`       //总供应
		PercentChange_1h  string    `json:"percent_change_1h"`  //近一小时价格变动
		PercentChange_24h string    `json:"percent_change_24h"` //
		PercentChange_7d  string    `json:"percent_change_7d"`  //
		LastUpdated       string    `json:"last_updated"`       //最近更新
	}

	var (
		err  error
		resp *http.Response
		body []byte
		data *ethData
		url  = "https://api.coinmarketcap.com/v1/ticker/ethereum/"
	)
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	str := strings.Trim(string(body), "[]")
	if err = json.Unmarshal([]byte(str), &data); err != nil {
		return
	}
	if _ret, _err := strconv.ParseFloat(data.PriceUsd, 64); _err != nil {
		return
	} else {
		fmt.Println(_ret)
	}
	fmt.Println(data.AvailableSupply)
	a, _ := json.Marshal(data)
	fmt.Println(string(a))
}
