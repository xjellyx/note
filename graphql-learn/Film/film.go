package main

import (
	"github.com/LnFen/handler"
	"github.com/graph-gophers/graphql-go"
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
type Resolver struct{}

func (r *Resolver) Film() *filmResolver {
	return &filmResolver{Hidden_Man}
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
	var schema = graphql.MustParseSchema(schemaStr, &Resolver{})
	h := handler.New(
		&handler.Config{
			Schema:   schema,
			GraphiQL: true,
			Pretty:   true,
		},
	)
	http.Handle("/graphql", h)
	http.ListenAndServe(":1234", nil)
}
