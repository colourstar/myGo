package main

import(
	"fmt"
)

type Bag struct{
	item []int
}

func (this *Bag)InsertBag()  {
	fmt.Println("InsertBag")
}

func main()  {
	bagInstanse := &Bag{}
	bagInstanse.InsertBag()
	fmt.Printf("%v\n",&bagInstanse)
}