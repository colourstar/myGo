package main

import(
	"errors"
	"fmt"
	"time"
)

func RPCClient(ch chan string, req string) (string,error)  {
	ch <- req

	select {
	case ack := <- ch:
		fmt.Printf("Client Recv Ack:%s\n",ack)
		return ack,nil
	case <- time.After(time.Second):
		return "",errors.New("Time out")
	}
}

func RPCServer(ch chan string)  {
	for {
		s := <-ch
		fmt.Printf("get client msg:%s\n",s)

		time.Sleep(time.Second * 2)

		ch <- "roger"
	}
}

func main()  {
	ch := make(chan string)
	go RPCServer(ch)

	recv , err := RPCClient(ch,"Hi Colourstar\n")

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("client recved:%s\n",recv)
	}
}