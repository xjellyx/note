package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-handler"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// Good 商品
type Good struct {
	ID    string  `json:"id"`    // id
	Name  string  `json:"name"`  // 名称
	Price float64 `json:"price"` // 价格
	Url   string  `json:"url"`   // url地址
}

var CacheMap = make(map[string]interface{})

// goodType 商品类型
var goodType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Good",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.String,
				Description: "id",
			}, "name": &graphql.Field{
				Type:        graphql.String,
				Description: "商品名称",
			}, "price": &graphql.Field{
				Description: "商品价格",
				Type:        graphql.Float,
			}, "url": &graphql.Field{
				Type:        graphql.String,
				Description: "商品url地址",
			},
		},
		Description: "商品参数类型",
	},
)

// goodListType 商品列表类型
var goodListType = graphql.NewList(goodType)

// queryType query 下的定义
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{ // 无需处理参数
			"goodList": &graphql.Field{
				Type: goodListType, // 处理结构体的回调函数，直接返回处理完成的结构体即可
				Resolve: func(p graphql.ResolveParams) (ret interface{}, err error) {
					var data []Good

					for _, v := range CacheMap {
						fmt.Printf("%T", v)
						switch _v := v.(type) {
						case *Good:
							data = append(data, *_v)
						}
					}

					ret = data
					return
				},
				Description: "获取商品列表",
			}, // 参数是id
			"good": &graphql.Field{
				Type: goodType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "商品id",
					},
				},
				Resolve: func(p graphql.ResolveParams) (ret interface{}, err error) { // 获取参数
					if _v, isOK := p.Args["id"].(string); isOK {
						data := CacheMap[_v]
						ret = data
						return
					} else {
						err = errors.New("Field 'goods' is missing required arguments: id. ")
						return
					}
				},
				Description: "商品",
			},
		},
	},
)

// goodInputType 商品输入参数类型定义
var goodInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "goodInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type:        graphql.String,
				Description: "名称",
			}, "price": &graphql.InputObjectFieldConfig{
				Description: "价格",
				Type:        graphql.Float,
			}, "url": &graphql.InputObjectFieldConfig{
				Type:        graphql.String,
				Description: "url",
			},
		},
		Description: "输入",
	},
)

// mutationType mutation
var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addGood": &graphql.Field{
				Description: "添加商品",
				Type:        goodType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type:        goodInputType,
						Description: "输入参数类型",
					},
				},
				Resolve: func(p graphql.ResolveParams) (ret interface{}, err error) {
					input, isOk := p.Args["input"].(map[string]interface{})
					if !isOk {
						err = errors.New("Field 'addGoods' is missing required arguments: input. ")
						return
					}
					var g = new(Good)
					// 生成一个id
					g.ID = uuid.NewV4().String()
					data, _err := json.Marshal(input)
					if _err != nil {
						err = _err
						return
					}
					if err = json.Unmarshal([]byte(data), &g); err != nil {
						return
					}

					// 用map保存数据
					CacheMap[g.ID] = g
					ret = g
					return
				},
			},
		},
	},
)

// schema 开机模式
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:        queryType,
		Mutation:     mutationType,
		Subscription: queryType,
	},
)

// Register 注册路由
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
	http.Handle("/", h)
	http.ListenAndServe("127.0.0.1:8081", nil)
}
