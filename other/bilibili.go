package main

import (
	"github.com/suboat/sorm"
	_ "github.com/suboat/sorm/driver/mongo"
)

func main() {

	conn := `{"url":"mongodb://127.0.0.1:27017/", "db": "business"}`
	if db, err := orm.New(orm.DriverNameMongo, conn); err != nil {
		panic(err)
	} else {
		defer db.Close()
	}
}
