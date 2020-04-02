package util

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"log"

	// 数据库驱动包
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql")
	driverName := beego.AppConfig.String("driverName")

	// 数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	// dbConn := "root:13572281710@tcp(127.0.0.1:3306)/testDemo?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		beego.Error(err.Error())
		return
	} else {
		db = db1
	}
}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// 例如：创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64)
		password VARCHAR(64)
		status INT(4)
		createtime INT(10)
		);`
	ModifyDB(sql)
}
