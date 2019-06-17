package main

import (
	"github.com/LnFen/graphql-api/ctrl"
	"github.com/LnFen/graphql-api/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type film struct {
	Name    string   //电影名
	Country string   //国家
	Stars   []string //主演
	Year    int32    //年份
	Runtime int32    //时长(分钟)
	Color   bool     //是否为彩色
	Score   float64  //评分
}

var schemaStr = `
		 schema {
            query: Query
        }
        type Query {
            film: Film
			all: [Film]!
}
        type Film {
            name   : String!
            country: String!
            year   : Int!
            runtime: Int!
            color  : Boolean
            score  : Float!
			stars : [String!]!
        }

`
var Hidden_Man = &film{
	Name:    "邪不压正",
	Stars:   []string{"姜文", "彭于晏", "廖凡", "周韵", "许晴"},
	Country: "China",
	Year:    2018,
	Runtime: 137,
	Color:   true,
	Score:   7.1,
}

type filmResolver struct {
	f *film
}

// NodeRoot 普通节点
type NodeRoot struct {
	ctrl.Node
	ctrl.NodeHandler
}

// Ctrl 是业务API
type Ctrl struct {
	ctrl.CTRL // 标准框架
}

var Root = &NodeRoot{}
var Api = &Ctrl{}

func (r *NodeRoot) Film() *filmResolver {
	return &filmResolver{Hidden_Man}
}
func (r *NodeRoot) All() (ret []*filmResolver) {
	var f *filmResolver
	f = new(filmResolver)
	f.f = &film{
		Name:    "邪不压正",
		Stars:   []string{"姜文", "彭于晏", "廖凡", "周韵", "许晴"},
		Country: "China",
		Year:    2018,
		Runtime: 137,
		Color:   true,
		Score:   7.1,
	}
	ret = append(ret, f)
	f.f = &film{
		Name:    "dasds",
		Stars:   []string{"姜文", "彭于晏", "廖凡", "周韵", "许晴"},
		Country: "sdasd",
		Year:    2018,
		Runtime: 137,
		Color:   true,
		Score:   7.1,
	}
	ret = append(ret, f)
	return
}
func (f *filmResolver) Name() string {
	return f.f.Name
}

func (f *filmResolver) Country() string {
	return f.f.Country
}

func (f *filmResolver) Year() int32 {
	return f.f.Year
}

func (f *filmResolver) Runtime() int32 {
	return f.f.Runtime
}

func (f *filmResolver) Color() *bool {
	return &f.f.Color
}
func (r *filmResolver) Stars() []string {
	return r.f.Stars
}
func (f *filmResolver) Score() float64 {
	return f.f.Score
}
func main() {
	var h = httprouter.New()
	node, _ := ctrl.NewNode(&Api.CTRL, schemaStr, Root)
	r, _ := router.NewRouter(nil)
	r.Prefix = "/api/test"
	r.ALL("/graphql", node.Handler)
	*r.CorsAllow = router.CorsAllowAll
	r.BindRouter(h)
	http.ListenAndServe(":1234", h)
}
