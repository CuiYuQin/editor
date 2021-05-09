package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//全局变量，数据库实例
var db *sql.DB

//数据库配置
const (
	host     = "127.0.0.1"
	port     = 3306
	user     = "root"
	password = "123456"
	dbname   = "ebook"
)

//初始化连接数据库
func Initdb() (err error) {

	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		user, password, host, port, dbname)
	db, err = sql.Open("mysql", dbInfo)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil

}

