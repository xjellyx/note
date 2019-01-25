package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	type exchangeRate struct {
		BankName string `json:"bankName"`
		Date     string `json:"date"`
		FSellPri string `json:"fSellPri"`
	}
	type rateData struct {
		Result []*exchangeRate
	}

	var (
		err error
		//美元实时汇率
		url      = "http://apicloud.mob.com/exchange/rmbquot/query?key=2970534a3dbfe&bank=1"
		resp     *http.Response
		respByte []byte
		data     *rateData
	)
	if resp, err = http.Get(url); err != nil {
		fmt.Println(err)
		return
	}
	if respByte, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println(err)
		return
	}
	str := string(respByte)
	fmt.Println(str)
	if err = json.Unmarshal([]byte(string(str)), &data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data.Result[0].FSellPri)
	for _, _v := range data.Result {
		fmt.Println(_v.FSellPri)
	}
	a := data.Result[0].FSellPri
	if _ret, _err := strconv.ParseFloat(a, 64); _err == nil {
		fmt.Println(_ret / 100)
	}
}
