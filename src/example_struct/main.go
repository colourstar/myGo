package main

import(
	"fmt"
)

type Point struct{
	x int
	y int
}

func main()  {
	// var point Point
	// point.x = 0
	// point.y = 1

	// point := new(Point)
	// point.x = 1
	// point.y = 2

	point := &Point{}
	point.x = 1
	point.y = 2
	
	fmt.Println("x = ",point.x,",y = ",point.y)
}