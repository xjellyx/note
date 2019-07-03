package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
)

func main()  {
	var fields=graphql.Fields{
		"hello":&graphql.Field{
			Type:graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				return "world",nil
			},
		},
	}
	rootQuery:=graphql.ObjectConfig{Name:"rootQuery",Fields:fields}
	schemaConfig:=graphql.SchemaConfig{
		Query:graphql.NewObject(rootQuery),
	}
	schem,err:=graphql.NewSchema(schemaConfig)
	if err!=nil{
		panic(err)
	}
	query:=`
		{
		hello
		}
	`
	params:=graphql.Params{Schema:schem,RequestString:query}
	r:=graphql.Do(params)
	if len(r.Errors)>0{
		panic(r.Errors)
	}
	rJSON,_:=json.Marshal(r)
	fmt.Println(string(rJSON))
}
