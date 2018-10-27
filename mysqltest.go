package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MMysql struct {
	conn_   *sql.DB
	lastsql string
	isclose bool

	dbname string
	dbport int
	dbuser string
	dbpass string
	dbhost string
}

func (op *MMysql) init() bool {
	connect_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", op.dbuser, op.dbpass, op.dbhost, op.dbport, op.dbname)
	conn, err := sql.Open("mysql", connect_str)
	if err == nil {
		op.conn_ = conn
		return true
	} else {
		return false
	}
}
func (op *MMysql) Init(name, user, pass, host string, port int) bool {
	op.dbhost = host
	op.dbname = name
	op.dbpass = pass
	op.dbport = port
	op.dbuser = user
	return op.init()
}
func (op *MMysql) QueryRecord(sqlstr string, query_results *[]map[string]string) int {
	var row_size = 0
	//println(sqlstr)
	rows, err := op.conn_.Query(sqlstr)
	if err == nil {
		defer rows.Close()
		columns, err := rows.Columns()
		if err != nil {
			return -1
		}
		values := make([]sql.RawBytes, len(columns))
		scanArags := make([]interface{}, len(values))
		for i := range values {
			scanArags[i] = &values[i]
		}
		for rows.Next() {
			row_size++
			err = rows.Scan(scanArags...)
			if err != nil {
				return -1
			}
			var value string
			item := map[string]string{}

			for i, col := range values {
				// Here we can check if the value is nil (NULL value)
				if col == nil {
					value = "NULL"
				} else {
					value = string(col)
				}
				item[columns[i]] = value
				//fmt.Println(columns[i], ": ", value)
			}
			*query_results = append(*query_results, item)
		}
	} else {
		panic(err)
	}
	return row_size
}

func (op *MMysql) Execute(sqlstr string) bool {
	_, err := op.conn_.Exec(sqlstr)
	if err != nil {
		return false
	}
	return true
}
func (op *MMysql) Close() {
	op.isclose = true
	op.conn_.Close()
}

func main() {
	fmt.Println("hello")
	var sql = MMysql{}
	sql.Init("t_test", "root", "passwd", "localhost", 3306)
	res := make([]map[string]string, 0)
	sql.QueryRecord("select keyword from t_whitelist", &res)
	fmt.Println(res)
}
