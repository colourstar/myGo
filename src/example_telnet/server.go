package main

import(
	"fmt"
	"net"
	"bufio"
	"strings"
)

type(
	ServerMgr struct{
		
	}
)

var(
	pkServerMgr *ServerMgr
)

func (this *ServerMgr) ServerStart(address string, exitCode chan int)  {
	listen,err := net.Listen("tcp",address)
	if err != nil {
		fmt.Println(err.Error())
		exitCode <- 1
	}

	fmt.Println("listen success,add:" + address)

	defer listen.Close()

	for {
		conn,err := listen.Accept()

		if err != nil {
			fmt.Println(err.Error())
			exitCode <- 1
		}

		go this.HandleSession(conn,exitCode)
	}
}

// 处理会话消息
func (this *ServerMgr)HandleSession(conn net.Conn,exitCode chan int)  {
	fmt.Println("Session Start")

	reader := bufio.NewReader(conn)

	for {
		str,err := reader.ReadString('\n')
		if err == nil {
			str = strings.TrimSpace(str)

			// 处理指令
			if this.ProcessMsg(str, exitCode) == false{
				conn.Close()
				break
			}

			conn.Write([]byte(str + "\r\n"))
		}else{
			fmt.Println("Seesion Error:" + err.Error())
			conn.Close()
			break
		}
	}
}

func (this *ServerMgr)ProcessMsg(str string,exitCode chan int) bool {
	if strings.HasPrefix( str, "@close" ){
		fmt.Println("Net Close")
		return false
	}else if strings.HasPrefix( str, "@shutdown" ){
		exitCode <- 0
		return false
	}

	fmt.Println(str)
	return true
}

func Inst() *ServerMgr {
	if pkServerMgr == nil{
		pkServerMgr = new(ServerMgr)
	}

	return pkServerMgr
}