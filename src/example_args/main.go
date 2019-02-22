package main

import(
	"fmt"
	"bytes"
)

func JoinStrings(sList ...string) string  {
	var buffer bytes.Buffer
	for _,s := range sList {
		buffer.WriteString( s )
	}
	return buffer.String()
}

func main()  {
	fmt.Println(JoinStrings("heelo",",123123","fdasfa\n"))
	fmt.Println(JoinStrings("heelo",",123123","fdasfa","132312312"))
}