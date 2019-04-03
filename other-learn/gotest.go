package main

import (
	"fmt"
	"github.com/suboat/sorm"
	_ "github.com/suboat/sorm/driver/mysql"
)

type configDB struct {
	DbName   string `json:"dbName"`   //
	User     string `json:"user"`     //
	Password string `json:"password"` //
	Host     string `json:"host"`     //
	Port     string `json:"port"`     //
}
type Student struct {
	StuID string `sorm:"size(36);primary" json:"stuId"`
	Name  string `sorm:"size(32);index" json:"name"`
	Age   int    `sorm:"size(23);index" json:"age"`
}

var modelStu orm.Model

func main() {
	var (
		db     orm.Database
		dataDB = &configDB{
			DbName:   "mysql",
			User:     "root",
			Password: "123456",
			Host:     "127.0.0.1",
			Port:     "3306",
		}
		err error
	)
	str := fmt.Sprintf(`{"user":"%s", "password": "%s", "host": "%s", "port": "%s", "dbname": "%s",
"sslmode": "disable","database":"test"}`, dataDB.User, dataDB.Password, dataDB.Host, dataDB.Port, dataDB.DbName)
	if db, err = orm.New("mysql", str); err != nil {
		panic(err)
	}
	modelStu = db.Model("student")

	if err = modelStu.Ensure(&Student{}); err != nil {
		panic(err)
	}
	/*if err = create(modelStu); err != nil {
		panic(err)
	}*/

	//app := iris.New()
	//mvc.Configure(app.Party("/root"), meMVC)
	//app.Run(iris.Addr(":8080"))
}

/*func meMVC(app *mvc.Application) {
	app.Handle(new(MyController))
}

type MyController struct{}

func (m *MyController) Get() (ret *Student) {
	data := new(Student)
	if err := modelStu.Objects().Filter(orm.M{
		"stuid": "54545647",
	}).One(data); err == nil {
		ret = new(Student)
		ret = data
	} else {
		panic(err)
	}
	return
}

func (m *MyController) create(id, name string, age int) (err error) {
	var stu *Student
	stu = new(Student)
	stu.StuID = id
	stu.Name = name
	stu.Age = age
	if err = modelStu.Objects().Create(stu); err != nil {
		return
	}
	return
}*/
