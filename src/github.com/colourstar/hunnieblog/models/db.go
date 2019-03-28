package models

import(
	"fmt"
	"time"
	"database/sql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/colourstar/hunnieblog/conf"
)

var (
	strDsn 			string
	pkSqlDataBase	*sql.DB
)

func InitDB(){
	orm.DefaultTimeLoc = time.UTC

	// // 连接DB
	// pkSqlDataBase = connectDB(conf.DBUser,conf.DBPassword,conf.DBHost,conf.DBPort,conf.DBName)
	// if ( pkSqlDataBase == nil ) {
	// 	beego.Trace("[db] : connectDB error")
	// }
}

// 创建table表
func CreateDB(){
	beego.Trace("[db] : CreateDB")

	var strCreateDBSql string

	switch conf.DBType {
	case "mysql":
		strCreateDBSql := fmt.Sprintf("CREATE DATABASE if not exist `%s` CHARSET utf8 COLLATE utf8_general_ci", conf.DBName)
		beego.Trace("[db] : sql is :", strCreateDBSql)
	default:
		beego.Critical("[db] : Error,not support :" , conf.DBType)
		return
	}

	pDBInst := connectDB()
	defer pDBInst.Close()

	_,err := pDBInst.Exec(strCreateDBSql)

	if (err != nil){
		beego.Error("[db] : db exec error,", err)
		panic(err.Error())
	}

	beego.Trace("[db] : DB Created : ", conf.DBName)
}

// 连接数据库,返回sql.DB的实例
func connectDB() (*sql.DB){
	strDsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8",conf.DBUser,conf.DBPassword,conf.DBHost,conf.DBPort)
	beego.Trace("[db] : connectDB,dsn =",strDsn)

	db,err := sql.Open("mysql",strDsn)

	if ( err != nil ){
		db.Close()
		panic(err)
	}

	return db
}

// 