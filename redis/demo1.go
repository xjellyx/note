package main

import (
	"github.com/go-redis/redis"
	"github.com/suboat/sorm/log"
	"time"
)

type C struct {
	Conn string
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Password: "123456",
		DB: 0,
	})
	log.Debug(client.Ping().Result())
	var (
		err error
	)
	c := new(C)
	c.Conn = "ssssssss"
	// 键值对十秒后过期
	if err = client.Set("key", []byte(c.Conn), time.Second*5).Err(); err != nil {
		log.Error(err)
	}

	d, _ := client.Get("key").Result()
	println("ssssssssssssss", d)
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
