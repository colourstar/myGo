package main

import (
	_ "github.com/colourstar/hunnieblog/conf"

	"github.com/colourstar/hunnieblog/models"
	"github.com/colourstar/hunnieblog/routers"

	"github.com/astaxie/beego"
)

func init(){
	beego.Trace("[main] : init")
	
	beego.Trace("[main] : models Run")
	models.Run()

	beego.Trace("[main] : routers Run")
	routers.Run()
}

func main(){
	beego.Run()
}