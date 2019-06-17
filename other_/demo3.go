package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type result struct {
	r   *http.Response
	err error
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{
		Transport: tr,
	}

	resultChan := make(chan result, 1)

	// 发起请求
	// req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	req, err := http.NewRequest("GET", "http://www.google.com", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		resp, err := client.Do(req)
		pack := result{
			r:   resp,
			err: err,
		}
		resultChan <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		er := <-resultChan
		fmt.Println("Timeout!", er.err)
	case res := <-resultChan:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Println(string(out))

	}
	return
}

func main() {
	process()
	ctx := context.WithValue(context.Background(), "id", 1314520)
	ctx = context.WithValue(ctx, "session", "hello world")
	process1(ctx)
}

func process1(ctx context.Context) {
	ret, ok := ctx.Value("id").(int)
	if !ok {
		ret = 7979 + 745
	}
	fmt.Println("ssssssssss", ret)

	s, _ := ctx.Value("session").(string)
	fmt.Println("aaaaaaaaa", s)
}
