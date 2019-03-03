package main

import(
	"fmt"
	"time"
)

func tickoutput()  {
	timecount := 0
	for  {
		timecount++
		fmt.Printf("time:%d\n",timecount)

		time.Sleep(time.Second)
	}
}

func main()  {
	go tickoutput()

	var scancontent string

	fmt.Scanln(&scancontent)
}