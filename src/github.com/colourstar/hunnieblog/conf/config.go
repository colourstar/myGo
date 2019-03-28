package conf

import(
	"fmt"
	"github.com/astaxie/beego"
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

func init(){
	Version := beego.AppConfig.DefaultString("version","0.0.1")
	beego.Trace(Version)

	DBHost = beego.AppConfig.String("DBHost")
	DBPort = beego.AppConfig.String("DBPort")
	DBUser = beego.AppConfig.String("DBUser")
	DBPassword = beego.AppConfig.String("DBPassword")
	DBName = beego.AppConfig.String("DBName")
	DBType = beego.AppConfig.String("DBType")

	beego.Trace(fmt.Sprintf("host:%s:%s,user:%s,password:%s,dbname:%s,dbtype:%s",DBHost,DBPort,DBUser,DBPassword,DBName,DBType))
}
