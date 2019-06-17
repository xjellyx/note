package obj

import (
	"database/sql"
	"fmt"
	"log"
)

type DB struct {
	db        *sql.DB
	TableName string
}

var (
	dbhost     = "127.0.0.1:3306"
	dbusername = "business"
	dbpassword = "business"
	dbname     = "business"
)

// Open
func (f *DB) Open() (err error) {
	var (
		db *sql.DB
	)
	if db, err = sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhost+")/"+dbname); err != nil {
		return
	}
	log.Println("打开数据库成功！！！")
	f.db = db
	return
}

// CLose
func (f *DB) Close() {
	defer f.db.Close()
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

func (f *DB) QueryFind(str string, args ...interface{}) (ret []dbRow, err error) {
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
func (f *DB) SelectAll(tableName string, args ...interface{}) (ret []dbRow, err error) {
	var (
		rows *sql.Rows
	)
	if rows, err = f.db.Query(fmt.Sprintf(`SELECT *FROM %s`, tableName), args...); err != nil {
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
func (f *DB) Column() (ret []dbRow, err error) {
	var (
		rows *sql.Rows
	)
	str := `SELECT count(*) FROM information_schema.tables WHERE table_name='%s'`
	str = fmt.Sprintf(str, f.TableName)
	if rows, err = f.db.Query(str); err != nil {
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
func (f *DB) Columns() (ret []dbRow, err error) {
	var columnLis = make([]dbRow, 0)
	str := `SELECT column_name FROM information_schema.columns WHERE table_name="%s"`
	str = fmt.Sprintf(str, f.TableName)
	if columnLis, err = f.QueryFind(str); err != nil {
		log.Println("select column ", err)
		return
	}
	ret = columnLis
	return
}
