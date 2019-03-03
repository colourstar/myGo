package main

import(
	"fmt"
	"sync"
)

var (
	iCount int
	mutexGuard sync.Mutex
)

func GetCount()int  {
	mutexGuard.Lock()

	defer mutexGuard.Unlock()
	return iCount
}

func SetCount(iValue int)  {
	mutexGuard.Lock()
	defer mutexGuard.Unlock()

	iCount = iValue
}

func main()  {
	SetCount(5)

	fmt.Printf("%d\n",GetCount())
}