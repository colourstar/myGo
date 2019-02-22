package main

import(
	"fmt"
)

func main()  {
	var cat int = 1
	var sss string = "ssssfdsafdf"
	var ptrsss = new(string)
	fmt.Printf("%p,%p,%p\n", &cat, &sss, ptrsss)
}