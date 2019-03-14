package main

import(
	"fmt"
	"os"
	"text/template"
)

type Latlng struct{
	Lat float32
	Lng float32
}

func (this Latlng)String() string  {
	return fmt.Sprintf("%g/%g",this.Lat, this.Lng)
}

func main()  {
	data := []template.FuncMap{}
	data = append(data,template.FuncMap{"name":"dotcool","url":"http://www.dotcoo.com","latlng":Latlng{24.2,135.1}})
	data = append(data,template.FuncMap{"name":"dotcool1","url":"http://www.dotcoo.com","latlng":Latlng{24.2,135.2}})
	data = append(data,template.FuncMap{"name":"dotcool2","url":"http://www.dotcoo.com","latlng":Latlng{24.2,135.3}})

	datatpl := `{{range .}}{{template "user" .}}{{end}}`
	usertpl := `{{define "user"}}name:{{.name}}, url:{{.url}}, latlng:{{.latlng}} lat:{{.latlng.Lat}} lng:{{.latlng.Lng}}{{end}}`

	tpl,err := template.New("data1").Parse(datatpl)

	if err != nil {
		panic(err)
	}

	_,err = tpl.Parse(usertpl)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	println()
}

