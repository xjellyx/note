package main

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func main() {
	var (
		client      = &http.Client{}
		req         *http.Request
		resp, resp1 *http.Response
		err         error
		doc         *goquery.Document
	)
	if req, err = http.NewRequest("GET", "https://eogdata.mines.edu/wwwdata/viirs_products/dnb_composites/v10/"+
		"/201806/vcmslcfg/SVDNB_npp_20180601-20180630_75N060W_vcmslcfg_v10_c201904251200.tgz", nil); err != nil {
		logrus.Fatalln(err)
		return
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Host", "eogdata.mines.edu")
	//req.Header.Set("Sec-Fetch-Site", "none")
	//req.Header.Set("Sec-Fetch-Dest", "document")
	//req.Header.Set("Sec-Fetch-Mode", "navigate")
	// req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	//req.Header.Set("Cookie", "mod_auth_openidc_state_V3FWHuc2ZbG1JFMIIvbonTTRiRQ=eyJhbGciOiAiZGlyIiwgImVuYyI6ICJBMjU2R0NNIn0..1QHxdjPSeQu4vZjc.dKLuojoNWDIqRwri93M_HiusIAnCyM6yWcrVfFahoYZy90Qq1gmsSVy3jHwlv6XqRXvfTDDgxUKog4GJjYqIqTsGnuzXj7FFHeuXB1c4Fb6ipsT2KjV5HkAniJiDcBLvjFwUIu71Il1ae8GbmE8TTsp6lE8IkBveMIpareceaO5VKltN--Altx3_un7N8cPvk3Oap4NHi2y1jYZnBAeDZRqiYWr-98G0kmAEfrT4Z0o5rKXMs6Ha8iY7qb8K92wHEAowP0zr4H5irPrSjZEVa-zi6Y1y8cBIuMDVp0isOLoCea5RkrnFnWmX0LEKEDjsptrGk55eGSAMzv4eAgI_9URQbjHZAadsM9Tq7pTsiRLvKA9CrbH2D58O6uwXSwxAXRBxLjVF3ZN3MY-8dTIhdS5qxtog_jzJgbCiNE06_XeoIJuU57UPvgj38n9S8k1CP5SE07dJzcboEKqSvARqpH9s_gnBD6FE7eRkc5LD-mfQ8kMQGiqzK2t457IJrvCGZhamqXl5PZQjt73Ed0XwOkYcJ7EOGLAoOurrsklMX-LkzvTqxdljaUNiyLZHlgy5GWaiuOeXBE7B8iun3Ph8y6wDmz-bXnnW5mI2cJQmjnpyNg.Fe8DJmQjokWKtmkK1-DYmQ")
	if resp, err = client.Do(req); err != nil {
		logrus.Fatalln(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		logrus.Fatalln(err)
		return
	}
	defer resp.Body.Close()
	if doc, err = goquery.NewDocumentFromReader(resp.Body); err != nil {
		logrus.Fatalln(err)
		return
	}
	cookie := resp.Header.Values("Set-Cookie")
	loginUrl, _ := doc.Find("#kc-form-login").Attr("action")
	// v中是登录帐号、密码等内容
	v := map[string]string{
		"username": "2685366884@qq.com",
		"password": "qwerdf123",
	}
	d, _ := json.Marshal(v)
	if req, err = http.NewRequest("POST", loginUrl, strings.NewReader(string(d))); err != nil {
		logrus.Fatalln(err)
		return
	}
	logrus.Warnln(loginUrl, strings.Join(cookie, ""))
	//req.Header.Add(":authority", "eogauth.mines.edu")
	//req.Header.Add(":method", "POST")
	//req.Header.Add(":scheme", "https")
	//req.Header.Add(":path", path)
	req.Header.Set("Cookie", strings.Join(cookie, ""))
	req.Header.Add("Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	req.Header.Add("Connection", "keep-alive")
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
	if resp1, err = client.Do(req); err != nil {
		logrus.Fatalln(err)
		return
	}
	logrus.Warnln(resp1.Header)
	if resp1.StatusCode != 200 {
		err = errors.New("got coolie failed")
		logrus.Fatalln(err)
		return
	}

}
