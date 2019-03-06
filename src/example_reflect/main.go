package main

import(
	"fmt"
	"reflect"
)

func main()  {
	var a = 0

	typeofA := reflect.TypeOf(a)

	fmt.Println(typeofA.Name(),typeofA.Kind())
}