package admin

import(
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            	int64
	Username		string 		`orm:"size(32)" form:"Username" valid:"Required;MaxSize(20);MinSize(6)"`
	Password      	string    	`orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
}

func (this *User) TableName() string{
	return "user_table"
}

func (this *User) Insert() error{
	if _, err := orm.NewOrm().Insert(this); err != nil{
		return err
	}
	return nil
}

func (this *User) Read(fields ...string) error{
	if err := orm.NewOrm().Read(this, fields...); err != nil{
		return err
	}
	return nil
}

func (this *User) Update(fields ...string) error{
	if _, err := orm.NewOrm().Update(this, fields...); err != nil{
		return err
	}
	return nil
}

func (this *User) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil{
		return err
	}
	return nil
}

func (this *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}
