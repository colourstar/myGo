package main

import(
	"fmt"
	// "reflect"
	// "sync"
	"container/list"
)

type Weapon int

const (
	Arrow Weapon = iota
	Shuriken
	SniperRifle
	Rifle
	Blower
)
func main()  {

	var arrayTest []string
	arrayTest = append(arrayTest, "1" )
	arrayTest = append(arrayTest, "heelo", "fdsadfasfd" )
	
	//fmt.Println(Arrow,Shuriken,SniperRifle,Rifle,Blower)
	//fmt.Println(reflect.TypeOf(SniperRifle))

	// for k,v := range arrayTest{
	// 	fmt.Println(k,v)
	// }
	// fmt.Println(arrayTest[1:2])
	// arrayTest = append(arrayTest,"123123")

	// var mapTest = make(map[string]int)
	// mapTest["route"] = 66
	// fmt.Println(mapTest["route"])
	// delete(mapTest,"route")
	// fmt.Println(mapTest["route"])

	// var syncMapTest sync.Map

	// syncMapTest.Store("a",1)
	// fmt.Println(syncMapTest.Load("a"))

	listTest := list.New()
	listTest.PushBack("abc")
	for i := listTest.Front(); i != nil; i = i.Next(){
		fmt.Println(i.Value)
	}
}