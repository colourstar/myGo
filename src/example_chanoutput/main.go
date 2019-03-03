package main

import(
	"fmt"
)

func printdata(c chan int) {
	for {
		s := <-c
		if (s == 0){
			break
		}

		fmt.Printf("%d\n",s)
	}
}

func main()  {
	c := make(chan int)

	go printdata(c)
	
	for int_i := 1; int_i < 10; int_i++ {
		c <- int_i
	}

	c <- 0
}