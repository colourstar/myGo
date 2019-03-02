package main

import(
	"fmt"
)

type DataWriter interface{
	Write(data interface{}) error
}

type File struct{

}

func (this *File) Write(data interface{})error  {
	fmt.Printf("Write %v\n",data)

	return nil
}

func main()  {
	var writer DataWriter

	file := &File{}

	writer = file

	writer.Write("Data")
}

