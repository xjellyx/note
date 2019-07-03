package obj

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
)

type DB struct {
	db        *sql.DB
	TableName string
}

var (
	dbhost     = "127.0.0.1:33306"
	dbusername = "business"
	dbpassword = "business"
	dbname     = "business"
)

// Open
func (d *DB) Open() (err error) {
	var (
		db *sql.DB
	)
	if db, err = sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhost+")/"+dbname); err != nil {
		return
	}
	log.Println("打开数据库成功！！！", dbusername+":"+dbpassword+"@tcp("+dbhost+")/"+dbname)
	d.db = db
	return
}

// Close
func (d *DB) Close() {
	defer d.db.Close()
}

//
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

// QueryFind 查看
func (d *DB) QueryFind(str string, args ...interface{}) (ret []dbRow, err error) {
	var (
		rows *sql.Rows
	)
	if rows, err = d.db.Query(str); err != nil {
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

// SelectAll 查看全部 TODO 这里供学习使用,实际项目不能这样写，数据量多的会影响性能
func (d *DB) SelectAll(tableName string, args ...interface{}) (ret []dbRow, err error) {
	var (
		rows *sql.Rows
	)
	if rows, err = d.db.Query(fmt.Sprintf(`SELECT *FROM %s`, tableName), args...); err != nil {
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

// Column
func (d *DB) Column() (ret []dbRow, err error) {
	var (
		rows *sql.Rows
	)
	str := `SELECT count(*) FROM information_schema.tables WHERE table_name='%s'`
	str = fmt.Sprintf(str, d.TableName)
	if rows, err = d.db.Query(str); err != nil {
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

// Columns
func (d *DB) Columns() (ret []dbRow, err error) {
	var columnLis = make([]dbRow, 0)
	str := `SELECT column_name FROM information_schema.columns WHERE table_name="%s"`
	str = fmt.Sprintf(str, d.TableName)
	if columnLis, err = d.QueryFind(str); err != nil {
		log.Println("select column ", err)
		return
	}
	ret = columnLis
	return
}

// Create 建表
func (d *DB) Create(sqlStr string) (ret *sql.Result, err error) {
	var (
		res sql.Result
	)
	if res, err = d.db.Exec(sqlStr); err != nil {
		return
	}

	ret = &res
	return
}

// ParamSQL 解析sql文件
func (d *DB) ParamSQL(filePath string) (ret string, err error) {
	var (
		data []byte
	)
	if data, err = ioutil.ReadFile(filePath); err != nil {
		return
	}

	ret = string(data)
	return
}

// Insert 插入数据
func (d *DB) Insert(sqlStr string) (err error) {
	if _, err = d.db.Exec(sqlStr); err != nil {
		return
	}

	return
}

// Update 更新数据
func (d *DB) Update(sqlStr string) (err error) {
	if _, err = d.db.Exec(sqlStr); err != nil {
		return
	}
	return
}
