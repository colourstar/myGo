package routers

import(
	"github.com/colourstar/hunnieblog/controllers/login"
	"github.com/astaxie/beego"
)

func init(){
	beego.Router("/",&controllers.LoginController{})
}