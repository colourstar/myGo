package main

import(
	"fmt"
	"sort"
)

type Sorter interface{
	Len() int

	Less(i,j int) bool

	Swap(i,j int)
}

type MyStringList []string

func (this MyStringList) Len()int  {
	return len(this)
}

func (this MyStringList) Less(i,j int) bool {
	return this[i] < this[j]
}

func (this MyStringList) Swap(i,j int){
	this[i], this[j] = this[j], this[i]
}

func main()  {
	testlist := MyStringList{
		"123123",
		"helo",
		"colourstar",
		"123321333123",
		"finaliy",
	}

	sort.Sort(testlist)

	for _,value := range testlist{
		fmt.Printf("%s\n",value)
	}
}