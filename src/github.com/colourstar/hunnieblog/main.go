package main

import (
	"flag"
	"os"
	"github.com/colourstar/hunnieblog/conf"				// config
	"github.com/colourstar/hunnieblog/models"			// models
	"github.com/colourstar/hunnieblog/routers"			// routers

	// third party
	"github.com/astaxie/beego"
)

func init(){
	beego.Trace("[main] : init")
	// help output
	bFlags_help := flag.Bool("h", false, "help")
	bFlags_createdb := flag.Bool("create", false, "create database")
	// flag parse
	flag.Parse()

	if (*bFlags_help){
		beego.Trace("-db	create database")
		beego.Trace("-clear	drop all database")
		os.Exit(0)
		return
	}
	// flag judge
	if (*bFlags_createdb){
		conf.Run()
		models.CreateDB()
		os.Exit(0)
		return
	}

	// conf run
	beego.Trace("[main] : config Run")
	conf.Run()
	// models run
	beego.Trace("[main] : models Run")
	models.Run()
	// routers run
	beego.Trace("[main] : routers Run")
	routers.Run()
	
}

func main(){
	// beego run
	beego.Trace("[main] : Run")

	
	beego.Run()
}