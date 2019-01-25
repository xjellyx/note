package obj

import (
	"database/sql"
	"fmt"
	"log"
)
type Mysql_db struct {
	db *sql.DB
	TableName string
}
var (
	dbhost     = "127.0.0.1:3306"
	dbusername = "root"
	dbpassword = "123456"
	dbname     = "mysql"
)

func (f *Mysql_db) MysqlOpen() {
	var (
		odb *sql.DB
		err error
	)
	str:=dbusername+":"+dbpassword+"@tcp("+dbhost+")/"+dbname
	println(str)
	if odb, err = sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhost+")/"+dbname); err != nil {
		fmt.Println(err, "打开数据库失败！！！")
	}
	fmt.Println("打开数据库成功！！！")
	f.db = odb
}
func (f *Mysql_db) MysqlClose() {
	defer f.db.Close()
	fmt.Println("数据库已经关闭！！！")
}
func scanRow(rows *sql.Rows) (dbRow, error) {
	columns, _ := rows.Columns()

	vals := make([]interface{}, len(columns))
	valsPtr := make([]interface{}, len(columns))

	for i := range vals {
		valsPtr[i] = &vals[i]
	}

	err := rows.Scan(valsPtr...)

	if err != nil {
		return nil, err
	}

	r := make(dbRow)

	for i, v := range columns {
		if va, ok := vals[i].([]byte); ok {
			r[v] = string(va)
		} else {
			r[v] = vals[i]
		}
	}

	return r, nil

}
type dbRow map[string]interface{}
func (f *Mysql_db) QueryFind(str string, args ...interface{}) (ret []dbRow, err error) {
	var (
		rows *sql.Rows
	)
	if rows, err = f.db.Query(str); err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		if r, _err := scanRow(rows); _err != nil {
			continue
		} else {
			ret = append(ret, r)
		}
	}
	return
}
func (f *Mysql_db) MysqlSelectAll(tableName string, args ...interface{}) (ret []dbRow, err error) {
	var (
		rows *sql.Rows
	)
	if rows, err = f.db.Query(fmt.Sprintf(`SELECT *FROM %s`, tableName), args...); err != nil {
		fmt.Println(err, "查询失败")
	}
	defer rows.Close()
	for rows.Next() {
		if r, _err := scanRow(rows); _err != nil {
			continue
		} else {
			ret = append(ret, r)
		}
	}
	return
}
func (f *Mysql_db)Clounm()(ret []dbRow,err error)  {
	var(
		rows *sql.Rows
	)
	str := `SELECT count(*) FROM information_schema.tables WHERE table_name='%s'`
	str = fmt.Sprintf(str, f.TableName)
	println(str)
	if rows,err=f.db.Query(str);err!=nil{
		return
	}
	defer rows.Close()
	for rows.Next(){
		if r, _err := scanRow(rows); _err != nil {
			continue
		} else {
			ret = append(ret, r)
		}
	}
	return
}
func (f *Mysql_db)Clounms()(ret []dbRow,err error)  {
	var columnLis = make([]dbRow, 0)
	str := `SELECT column_name FROM information_schema.columns WHERE table_name="%s"`
	str = fmt.Sprintf(str, f.TableName)
	if columnLis, err = f.QueryFind(str); err != nil {
		log.Println("select column ", err)
		return
	}
	ret=columnLis
	return
}