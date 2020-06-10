package main

import (
	"fmt"
	"github.com/olongfen/note/log"
	"github.com/robfig/cron/v3"
)

func main()  {
	initCron()
	select {

	}

}

func initCron()  {
	var (
	 secondParser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	c = cron.New(cron.WithParser(secondParser), cron.WithChain())

	)
	go func() {
		if _,err :=  c.AddFunc("*/5 * * * * ?", func() {
			fmt.Println("qqqqqqqqqqqqqqq")
		});err!=nil{
			log.Errorln(err)
			return
		}
	}()

	go func() {
		if _,err :=  c.AddFunc("*/5 * * * * ?", func() {
			fmt.Println("aaaaaaaaaaaaaaaaaaaaaa")
		});err!=nil{
			log.Errorln(err)
			return
		}
	}()

	c.Start()
}
