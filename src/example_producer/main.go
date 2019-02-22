package main

import(
	"fmt"
	"time"
)

func produce(content string, channel chan<- string){
	for {
		channel <- fmt.Sprintf("%s:%v",content,1)
		time.Sleep(time.Second)
	}
}

func consume(channel <-chan string){
	for {
		message := <-channel
		fmt.Printf(message)
	}
}

func main(){
	channel := make(chan string)
	go produce("produce",channel)
	consume(channel)
}