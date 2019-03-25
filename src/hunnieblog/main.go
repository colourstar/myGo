package main

import (
	_ "hunnieblog/conf"

	_ "hunnieblog/routers"
	"github.com/astaxie/beego"
)

func main(){
	beego.Run()
}