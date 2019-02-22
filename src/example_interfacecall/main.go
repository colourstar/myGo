package main

import(
	"fmt"
)

// 调用器的接口
type Invoker interface{
	Call(interface{})
}

// 结构体的类型
type Struct struct{

}

// 实现invoker的call
func (s *Struct) Call(p interface{}){
	fmt.Println("from struct", p)
}

// 函数定义为类型
type FuncCall func(interface{})

// 实现invoker的call
func (f FuncCall) Call(p interface{}){
	// 调用本体
	f(p)
}

func main()  {
	var invoker Invoker

	// 实例化
	s := new(Struct)

	invoker = s

	invoker.Call("hello")

	invoker = FuncCall(func(p interface{}){
		fmt.Println("from func caller", p)
	})

	invoker.Call("hello")
}