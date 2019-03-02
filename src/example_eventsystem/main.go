package main

import(
	"fmt"
)

type EventSystem struct{
	eventListMap map[string] []func(interface{})
}

func (this *EventSystem) RegisterCallback(evname string,callback func(interface{})){
	eventList := this.eventListMap[evname]
	eventList = append(eventList,callback)
	this.eventListMap[evname] = eventList
}

func (this *EventSystem) TrigEvent(evname string, param interface{}){
	eventList := this.eventListMap[evname]
	for _,callback := range eventList {
		callback(param)
	}
}

func GlobalEvent(param interface{}){
	fmt.Printf("trigger event,%d",param)
}


func main()  {
	evSystem := &EventSystem{}
	evSystem.eventListMap = make(map[string] []func(interface{}))
	evSystem.RegisterCallback("event",func(param interface{}){
		fmt.Printf("event trigger \n")
	})

	evSystem.TrigEvent("event",100)
}