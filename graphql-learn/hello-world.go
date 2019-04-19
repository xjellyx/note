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
	h := handler.HttpHandler(s, true, true)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8899", nil)
}
