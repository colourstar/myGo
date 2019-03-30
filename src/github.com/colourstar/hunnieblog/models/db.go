package models

import(
	"fmt"
	"time"
	"os"
	"database/sql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/colourstar/hunnieblog/conf"
)

var (
	strDsn 			string
	strOrmDsn		string
	pkSqlDataBase	*sql.DB
)

func InitDB(){
	orm.DefaultTimeLoc = time.UTC
}

// 创建table表
func CreateDB(){
	beego.Trace("[db] : CreateDB")

	var strCreateDBSql string

	switch conf.DBType {
	case "mysql":
		strCreateDBSql = fmt.Sprintf("CREATE DATABASE if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", conf.DBName)
	default:
		beego.Critical("[db] : Error,not support :" , conf.DBType)
		return
	}

	pDBInst := connectDB()
	defer pDBInst.Close()

	beego.Trace("[db] :", strCreateDBSql)
	_,err := pDBInst.Exec(strCreateDBSql)

	if (err != nil){
		beego.Error("[db] : db exec error,", err)
		panic(err.Error())
	}

	beego.Trace("[db] : DB Created : ", conf.DBName)

	registerORM()
}

// 删除数据库
func DropDB(){
	beego.Trace("[db] : DropDB")

	var strDropDBSql string

	switch conf.DBType {
	case "mysql":
		strDropDBSql = fmt.Sprintf("DROP DATABASE if exists `%s`", conf.DBName)
	default:
		beego.Critical("[db] : Error,not support :" , conf.DBType)
		return
	}

	pDBInst := connectDB()
	defer pDBInst.Close()

	beego.Trace("[db] :", strDropDBSql)
	_,err := pDBInst.Exec(strDropDBSql)

	if (err != nil){
		beego.Error("[db] : db exec error,", err)
		panic(err.Error())
	}

	beego.Trace("[db] : DB droped : ", conf.DBName)
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

// 注册orm
func registerORM(){
	switch conf.DBType {
	case "mysql":
		strOrmDsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",conf.DBUser,conf.DBPassword,conf.DBHost,conf.DBPort,conf.DBName)
		orm.RegisterDriver("mysql",orm.DRMySQL)
		break
	default:
		beego.Critical("[db] : registerORM Error,not support :" , conf.DBType)
		return
	}

	// register db
	beego.Trace("[db] : registerORM,dsn = ", strOrmDsn)
	err := orm.RegisterDataBase("default",conf.DBType, strOrmDsn)
	if err != nil {
		beego.Error("[db] : register error:" + err.Error())
		panic(err.Error())
	}

	// DBLog
	if (conf.DBLog == "open"){
		beego.Trace("[db] : registerORM open dblog")
		orm.Debug = true
		w, e := os.OpenFile("log/db.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if (e != nil) {
			beego.Error("[db] : open file failed")
			panic(err.Error())
		}
		orm.DebugLog = orm.NewLog(w)
	}

	// db model
	RegisterModel()

	// db syncdb
	err = orm.RunSyncdb("default",true,true)
	if err != nil {
		beego.Error("[db] : database sync error:" + err.Error())
		panic(err.Error())
	}
}

// 