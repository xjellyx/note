package main

import (
	"fmt"
)

func main() {
	str := `
schema {
		query: Query
	}
	# 获取信息
	type Query{
		getInfo(id:String!):String!
	}
`
	a, _err := graphql.ParseSchema(str, "")
	fmt.Println(a, _err)
}
