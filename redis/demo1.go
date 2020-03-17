package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/suboat/sorm/log"
	"sync"
	"time"
)

type C struct {
	Conn string
	Data *sync.Map `json:"-"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "business",
		DB:       0,
	})
	log.Debug(client.Ping().Result())
	var (
		err error
	)
	c := new(C)
	c.Conn = "ssssssss"
	c.Data = new(sync.Map)

	// 键值对十秒后过期
	if err = client.Set("key", []byte(c.Conn), time.Second*5).Err(); err != nil {
		log.Error(err)
	}
	time.Sleep(time.Second * 6)
	_d, _e := client.Exists("key").Result()
	fmt.Println(_d, "aaaaaaaaaaaa", _e)
	c.Data.Store("uid", "username")
	_dd, _ := json.Marshal(c)
	fmt.Println(_dd)
	client.HSet("room", "demo1", "qqqqqq")
	client.HSet("room", "demo2", "dqeqweqweqwe")
	client.HSet("room", "demo1", "qweqwdgtryhtu")
	client.HSet("room", "demo3", string(_dd))
	d := client.HGet("room", "demo3").Val()

	fmt.Println(d)
	fmt.Println(client.HGetAll("room"))
	//// 延迟8秒获取
	//time.Sleep(time.Second * 3)
	//if val, _err := client.Get("key").Result(); _err != nil {
	//	log.Error(_err)
	//} else {
	//	fmt.Println(val)
	//}
	//// 再延迟3秒获取，会报错
	//time.Sleep(time.Second * 3)
	//if val, _err := client.Get("key").Result(); _err != nil {
	//	log.Error(_err)
	//} else {
	//	fmt.Println(val)
	//}
	//
	//// SET key value EX 10 NX
	//set, err := client.SetNX("key", "value", 100*time.Second).Result()
	//
	//fmt.Println(set)

}
