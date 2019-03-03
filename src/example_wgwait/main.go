package main

import(
	"fmt"
	"net/http"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	urlList := []string{
		"https://www.baidu.com",
		"https://www.qq.com",
		"https://www.sina.com.cn",
	}

	for _,url := range urlList {
		wg.Add(1)

		go func(url string){
			defer wg.Done()

			_,err := http.Get(url)

			fmt.Println(url,err)

		}(url)
	}

	wg.Wait()

	fmt.Println("url over")
}