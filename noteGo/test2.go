package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/suboat/sorm"
	_ "github.com/suboat/sorm/driver/mysql"
)

type User struct {
	UID      string `sorm:"size(36); primary" json:"uid"`
	UserName string `sorm:"size(16); index" json:"userName"`
	Password string `sorm:"size(16); index" json:"password"`
	Name     string `sorm:"size(16); index" json:"name"`
	Age      int32  `sorm:"index" json:"age"`
	Phone    string `sorm:"size(16); index" json:"phone"`
	Email    string `sorm:"size(16); index" json:"email"`
	Sex      string `sorm:"size(5); index" json:"sex"`
}

func main() {
	var (
		db  orm.Database
		err error
	)
	u := uuid.NewV5(uuid.UUID{}, "nanghjy")
	fmt.Println(u)
	str := `{"user":"root", "password": "123456", "host": "127.0.0.1", "port": "3306", 
"sslmode": "disable","database":"project"}`
	if db, err = orm.New(orm.DriverNameMysql, str); err != nil {
		return
	}
	if err = db.Model("people").Ensure(&User{}); err != nil {
		panic(err)
	}
}
