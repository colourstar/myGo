package conf

import(
	"fmt"
	"github.com/astaxie/beego"
)

// 参数数据库控制
type 	FlagEventOp 	int
const(
	_ FlagEventOp = iota
	CreateDB					// 创建数据库
	DeleteDB					// 删除数据库
	truncateDB					// 清空所有数据
)

var (
	Version		string

	// DataBase
	DBHost		string
	DBPort		string
	DBUser		string
	DBPassword	string
	DBName		string
	DBType		string
)

type FlagConfig struct{
	DBOpType	*string
	DBHelp		*bool
}

func Run(){
	Version := beego.AppConfig.DefaultString("version","0.0.1")
	beego.Trace(Version)

	DBHost = beego.AppConfig.String("DBHost")
	DBPort = beego.AppConfig.String("DBPort")
	DBUser = beego.AppConfig.String("DBUser")
	DBPassword = beego.AppConfig.String("DBPassword")
	DBName = beego.AppConfig.String("DBName")
	DBType = beego.AppConfig.String("DBType")

	beego.Trace(fmt.Sprintf("[conf] : DBConf -> host:%s:%s,user:%s,password:%s,dbname:%s,dbtype:%s",DBHost,DBPort,DBUser,DBPassword,DBName,DBType))
}
