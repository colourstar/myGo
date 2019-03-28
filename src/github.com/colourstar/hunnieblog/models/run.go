package models

import(
	"github.com/astaxie/beego"
)

func Run(){
	// 初始化数据库
	beego.Trace("[model] : Model Run Start")
	InitDB()
}