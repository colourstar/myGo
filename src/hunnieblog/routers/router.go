package routers

import(
	"hunnieblog/controllers/login"
	"github.com/astaxie/beego"
)

func init(){
	beego.Router("/",&controllers.LoginController{})
}