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

	// 连接DB
	pkSqlDataBase = connectDB(conf.DBUser,conf.DBPassword,conf.DBHost,conf.DBPort,conf.DBName)
	if ( pkSqlDataBase == nil ) {
		beego.Trace("[db] : connectDB error")
	}
}

// 连接数据库
func connectDB(strUser string,strPasswd string,strAddr string,strPort string,strDBName string) (*sql.DB){
	strDsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",strUser,strPasswd,strAddr,strPort,strDBName)
	beego.Trace("[db] : connectDB,dsn =",strDsn)

	db,err := sql.Open("mysql",strDsn)

	if ( err != nil ){
		db.Close()
		panic(err)
		return nil
	}

	return db
}

// 创建table
func Createtable(){
	beego.Trace("[dbhelper] : data init start")

	beego.Trace("[dbhelper] : data init end")
}

// 