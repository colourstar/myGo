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
	DBHost		string		// host
	DBPort		string		// port
	DBUser		string		// user
	DBPassword	string		// passwd
	DBName		string		// DBName
	DBType		string		// DB类型,目前只支持mysql
	DBLog		string 		// 开启DB日志
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
	DBLog = beego.AppConfig.String("DBLog")

	beego.Trace(fmt.Sprintf("[conf] : DBConf -> host:%s:%s,user:%s,password:%s,dbname:%s,dbtype:%s",DBHost,DBPort,DBUser,DBPassword,DBName,DBType))
}
