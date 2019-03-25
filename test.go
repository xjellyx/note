package main

import (
	"fmt"
	"regexp"
)

func main() {
	var (
		str = "{\"HundredPercent\":100,\"BaseExchangeCashRate\":100,\"ShareExchangeCashRate\":10," +
			"\"ZPayExchangeCashRate\":500,\"CashDiscountRate\":85," +
			"\"BaseExchangeEthRate\":1000,\"Poundage\":50000000,\"SellBasePoundage1000\":100000000,\"SellBasePoundage5000\":300000000,\"SellBasePoundage10000\":500000000,\"SellBasePoundage30000\":1000000000,\"SellBasePointT1\":30000,\"SellBasePointT2\":100001,\"SellBasePointT3\":500000,\"SellBasePointT4\":1000000,\"SellBaseMonthQuotaT1\":4501,\"SellBaseMonthQuotaT2\":8000,\"SellBaseMonthQuotaT3\":30000,\"SellBaseMonthQuotaT4\":45000,\"SystemCash\":\"CNY\",\"Power\":100000000,\"DefaultDecimal\":8,\"CfgAutoRound\":false,\"SellBaseCashPointMin\":0,\"BuyBaseQuotaTimes\":1.6,\"SwitchBuyTrade\":true}"
	)
	fmt.Println(str)
	reg := regexp.MustCompile(`"BaseExchangeEthRate":.*?,`)
	s := reg.Find([]byte(str))
	fmt.Println(string(s))
	a := reg.ReplaceAllString(str, `${1}"BaseExchangeEthRate":58888,`)
	fmt.Println(a)

}

func getSetting(token string) (ret string, err error) {
	str := fmt.Sprintf(`query setingget{
  me(token:%s){
    getSettingJson
  }
}`, token)

	return
}
