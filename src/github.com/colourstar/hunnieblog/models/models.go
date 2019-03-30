package models

import(
	"github.com/astaxie/beego/orm"

	"github.com/colourstar/hunnieblog/models/admin"
)

func RegisterModel(){
	orm.RegisterModel(new(admin.User))
}

