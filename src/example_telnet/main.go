package main

import(
	"os"
	"fmt"
)

func main()  {
	fmt.Println("Main Start!!!")

	exitCode := make(chan int)

	serverMgr := Inst()
	go serverMgr.ServerStart("127.0.0.1:7001",exitCode)

	code := <-exitCode

	fmt.Printf("Server Exit!,Exit Code:%d",code)

	os.Exit(code)
}

