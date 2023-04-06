package main

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("template.txt")
	str := string(data)
	titleAllReg := regexp2.MustCompile("(?<=风险预警：[\\n]+).*(?=[\\n]+预警)", 0)
	m, _ := titleAllReg.FindStringMatch(str)
	countryReg := regexp2.MustCompile("(?<=\\[).*(?=\\])", 0)
	country, _ := countryReg.FindStringMatch(m.String())
	fmt.Println(country.String())
	titleReg := regexp2.MustCompile("(?<=\\][\\s]+).*", 0)
	title, _ := titleReg.FindStringMatch(m.String())
	fmt.Println(title)
	waringReg := regexp2.MustCompile("(?<=预警.[\\n]*\\[级别.).*(?=\\][\\n]*)", 0)
	waring, _ := waringReg.FindStringMatch(str)
	fmt.Println(waring)
	dateReg := regexp2.MustCompile("\\d{4}-\\d{1,2}-\\d{1,2}", 0)
	date, _ := dateReg.FindStringMatch(str)
	fmt.Println(date)
	detailReg := regexp2.MustCompile("(?<=风险描述：[\\n]+).*(?=[\\n]+专家观点)", 0)
	detail, _ := detailReg.FindStringMatch(str)
	fmt.Println(detail)
	lonReg := regexp2.MustCompile("(?<=经度：).*(?=[\\n])", 0)
	lon, _ := lonReg.FindStringMatch(str)
	fmt.Println(lon)
	latReg := regexp2.MustCompile("(?<=纬度：).*(?=[\\n])", 0)
	lat, _ := latReg.FindStringMatch(str)
	fmt.Println(lat)
}
