package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/suboat/sorm"
	"github.com/suboat/sorm/driver/mysql"
	_ "github.com/suboat/sorm/driver/mysql"
)

type Config struct {
	TableAdmin string
	MainDb     *ConfigDb
}

// ConfigDb 数据库链接
type ConfigDb struct {
	DbName   string `json:"dbName"`   //
	User     string `json:"user"`     //
	Password string `json:"password"` //
	Host     string `json:"host"`     //
	Port     string `json:"port"`     //
}

func main() {
	var (
		config   *Config
		db       orm.Database
		mainConn string
		err      error
	)
	config = &Config{
		TableAdmin: "projectadmin",
		MainDb: &ConfigDb{
			DbName:   "mysql",
			User:     "root",
			Password: "123456",
			Host:     "127.0.0.1",
			Port:     "3306",
		},
	}
	mysql.CfgDbUnsafe = true
	mainConn = fmt.Sprintf(`{"user":"%s", "password": "%s", "host": "%s", "port": "%s", "dbname": "%s", 
"sslmode": "disable","database":"project"}`, config.MainDb.User, config.MainDb.Password, config.MainDb.Host,
		config.MainDb.Port, config.MainDb.DbName,
	)
	if db, err = orm.New(orm.DriverNameMysql, mainConn); err != nil {
		panic(err)
	}
	if err = model.Init(db.Model(config.TableAdmin)); err != nil {
		panic(err)
	}
	app := iris.Default()
	if err = ctrl.InitRouter(app); err != nil {
		panic(err)
	}
	if err = app.Run(iris.Addr(":8865")); err != nil {
		panic(err)
	}
}
