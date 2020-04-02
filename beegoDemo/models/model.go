package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// 数据库驱动包
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	driverName := beego.AppConfig.String("driverName")
	// 注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	// 数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	// dbConn := "root:13572281710@tcp(127.0.0.1:3306)/testDemo?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		beego.Error("连接数据库出错")
		return
	}
	beego.Info("连接数据库成功")
}
