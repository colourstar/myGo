package class1

import(
	"example_classfactory/base"
	"fmt"
)

type Class2 struct{

}

func (this *Class2)Do()  {
	fmt.Println("class2 do")
}

func init()  {
	base.GetInst().RegClass("class2",func () base.Classer  {
		return new(Class2)
	})
}

