package main

import(
	"fmt"
)

type DumStruct struct{
	a int
}

func dummy(iTest int) int {
	var iRet = 0
	iRet = iTest

	return iRet
}

func dummy1() *DumStruct{
	var a DumStruct
	return &a
}

func main()  {
	var iTest1 = 0
	fmt.Printf("%d,%d,%p\n", iTest1, dummy(1), dummy1())
}