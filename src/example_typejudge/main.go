package main

import(
	"fmt"
)

func typeoutput(data interface{})  {
	switch data.(type) {
	case int:
		fmt.Println("this is int\n")
	case string:
		fmt.Println("this is string\n")
	case float32:
		fmt.Println("this is float32\n")
	case float64:
		fmt.Println("this is float64\n")
	case bool:
		fmt.Println("this is bool\n")
	default:
		fmt.Println("this is nothing\n")
	}
}

func main()  {
	typeoutput(123)
	typeoutput("fdsafd")
	typeoutput(10.0)
	typeoutput(true)
}