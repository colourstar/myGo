package main

import(
	"example_classfactory/base"
	_ "example_classfactory/class1"
	_ "example_classfactory/class2"
)

func main()  {
	c1 := base.GetInst().CreateClass("class1")
	c2 := base.GetInst().CreateClass("class2")
	c1.Do()
	c2.Do()
}