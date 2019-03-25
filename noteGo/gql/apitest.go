package main

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-handler"
	"net/http"
	"strconv"
)

type Goods struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Url   string  `json:"url"`
}

var goodsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Goods",
		Fields: graphql.Fields{"id": &graphql.Field{
			Type: graphql.String,
		}, "name": &graphql.Field{
			Type: graphql.String,
		}, "price": &graphql.Field{
			Type: graphql.Float,
		}, "url": &graphql.Field{
			Type: graphql.String,
		},
		},
	},
)
var goodsListType = graphql.NewList(goodsType)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{ // 无需处理参数
			"goodsList": &graphql.Field{
				Type: goodsListType, // 处理结构体的回调函数，直接返回处理完成的结构体即可
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "sdasdasfbhddashfa", nil
				},
			}, // 参数是id
			"goods": &graphql.Field{
				Type: goodsType,
				Args: graphql.FieldConfigArgument{"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) { // 获取参数
					if _, isOK := p.Args["id"].(string); isOK {
						return "sdasdasfbhddashfa", nil
					}
					err := errors.New("Field 'goods' is missing required arguments: id. ")
					return nil, err
				},
			},
		},
	},
)
var goodsInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "goodsInput",
		Fields: graphql.InputObjectConfigFieldMap{"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		}, "price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		}, "url": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{"addGoods": &graphql.Field{
			Type: goodsType,
			Args: graphql.FieldConfigArgument{"input": &graphql.ArgumentConfig{
				Type: goodsInputType,
			},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				input, isOk := p.Args["input"].(map[string]string)
				if !isOk {
					err := errors.New("Field 'addGoods' is missing required arguments: input. ")
					return nil, err
				}
				a, err := strconv.ParseFloat(input["price"], 64)
				result := Goods{
					Name:  input["name"],
					Price: a,
					Url:   input["url"],
				} // 处理数据
				return result, err
			},
		},
		},
	},
)
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func Register() *handler.Handler {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return h
}

func main() {
	h := Register()
	http.Handle("/graphql", h)
	http.ListenAndServe("127.0.0.1:8081", nil)
}
