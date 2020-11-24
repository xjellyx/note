package data

import (
	"github.com/olongfen/note/data/mode/database"
	"github.com/olongfen/note/data/mode/http_data"
	"gorm.io/driver/postgres"
)

type Config struct {
	Address string
	Mode    string
}

func Open(c Config) (ret Object, err error) {
	switch c.Mode {
	case "http", "https":
		ret = http_data.NewDataHttp(c.Address)
	case "database":
		return database.NewDatabase(postgres.Open(c.Address), nil)
	case "ftp":
	case "websocket":

	}
	return
}

type Object interface {
	GetData() (interface{}, error)
}
