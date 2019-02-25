package main

import(
	"fmt"
	"runtime"
	"runtime/debug"
)

type panicContext struct{
	function string
}

func ProtectRun(entry func())  {
	// 延迟处理
	defer func(){
		// 获取宕机上下文
		err := recover()

		switch err.( type ) {
		case runtime.Error:
			fmt.Println("runtime error:", err )
		default:
			fmt.Println("error:", err )
		}

		fmt.Printf( string( debug.Stack() ) )
	}()

	entry()
}

func main()  {

	fmt.Println("Run")
	// ProtectRun(func(){
	// 	fmt.Println("手动宕机前")

	// 	// panic传递上下文
	// 	panic(&panicContext{
	// 		"手动触发panic",
	// 	})

	// 	fmt.Println("手动宕机后")
	// })

	// 故意造成空指针访问错误
	ProtectRun(func(){
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})

	fmt.Println("Run Over")

	// defer fmt.Println("111111111")
	// panic("crash")
}