package util

import(
	"fmt"
	"os"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
)

var(
	DBInst

)

func init(){
	orm.DefaultTimeLoc = time.UTC
}

// 连接数据库
func connectDB(){

}

// 创建table
func Createtable(){
	beego.Trace("[dbhelper] : data init start")

	beego.Trace("[dbhelper] : data init end")
}

// 