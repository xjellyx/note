package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type conf struct {
	User user
}

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	viper.SetConfigFile("demo.yaml")
	viper.SetDefault("user", user{
		Name: "Tom",
		Age:  "29",
	})
	c := &conf{}
	viper.WriteConfig()
	viper.ReadInConfig()
	if err := viper.Unmarshal(c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(viper.Get("user"))
	fmt.Println(c)
}
