package main

import (
	"github.com/LnFen/handler"
	"github.com/graph-gophers/graphql-go"
	"net/http"
)

var sch = `schema{
	query:Query
}
	type Query{
	hello: String!	
}
`

type query struct {
}

func (q *query) Hello() string {
	return "hello world ! ! !"
}

func main() {
	var s = graphql.MustParseSchema(sch, &query{})
	h := handler.New(
		&handler.Config{
			Schema:   s,
			Pretty:   true,
			GraphiQL: true},
	)
	http.Handle("/", h)
	http.ListenAndServe(":8899", nil)
}
