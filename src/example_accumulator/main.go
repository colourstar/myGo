package main

import(
	"fmt"
)

func Accumulator(iValue int) func()int {
	return func() int {
		iValue++
		return iValue
	}
}
func main()  {
	fAccumulator := Accumulator(1)

	fmt.Println(fAccumulator())
	fmt.Println(fAccumulator())

	fmt.Printf("%p\n",fAccumulator)

	fAccumulator2 := Accumulator(10)

	fmt.Println(fAccumulator2())
	fmt.Println(fAccumulator2())

	fmt.Printf("%p\n",fAccumulator2)
	fmt.Println(fAccumulator())
}