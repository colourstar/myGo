package main

import (
	_ "github.com/colourstar/hunnieblog/conf"
	_ "github.com/colourstar/hunnieblog/routers"

	"github.com/astaxie/beego"
)

func main(){
	beego.Run()
}