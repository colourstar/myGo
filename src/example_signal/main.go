package main

import(
	"fmt"
	"os"
	"os/signal"
)

func main()  {
	c := make( chan os.Signal,0)
	// 监听操作系统的打断消息
	signal.Notify(c, os.Interrupt, os.Kill)

	// 这里阻塞
	s := <-c
	fmt.Printf("client exit --> signal[%v]", s)
}