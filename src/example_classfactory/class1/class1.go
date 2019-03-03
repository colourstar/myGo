package class1

import(
	"example_classfactory/base"
	"fmt"
)

type Class1 struct{

}

func (this *Class1)Do()  {
	fmt.Println("class1 do")
}

func init()  {
	base.GetInst().RegClass("class1",func () base.Classer  {
		return new(Class1)
	})
}

