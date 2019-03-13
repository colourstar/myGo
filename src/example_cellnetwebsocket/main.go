package main

import(
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/golog"

	"fmt"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/json"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
	"reflect"
)

var log = golog.New("websocket_server")

type TestEchoAct struct{
	Msg 	string
	Value 	int32
}

func (self *TestEchoAct)String() string{
	return fmt.Sprintf("%+v", *self)
}

// 注册消息
func init()  {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec : codec.MustGetCodec("json"),
		Type : reflect.TypeOf((*TestEchoAct)(nil)).Elem(),
		ID : 1234,
	})
}

// 运行服务器
func main()  {
	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("gorillaws.Acceptor","server","http://127.0.0.1:18802~18803/echo",queue)

	proc.BindProcessorHandler(p,"gorillaws.ltv",func(ev cellnet.Event){
		switch message:= ev.Message().(type) {
		case *cellnet.SessionAccepted:
			log.Debugln("server accepted")
		case *cellnet.SessionConnected:
			log.Debugln("server Connected")
		case *cellnet.SessionClosed:
			log.Debugln("server Closed")
		case *TestEchoAct:
			log.Debugf("recv: %+v", message)

			ev.Session().Send(&TestEchoAct{
				Msg : "roger",
				Value : 1024,
			})
		}
	})

	// 开始监听
	p.Start()

	queue.StartLoop()

	queue.Wait()
}