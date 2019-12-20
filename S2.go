package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"log"
)

//queue
func main() {
	//父亲
	C := colly.NewCollector(
		colly.AllowedDomains(""),
	)
	//儿子
	url := ""
	queue2, err := queue.New(
		2, // consumer threads
		&queue.InMemoryQueueStorage{
			MaxSize: 10000, // queue2 storage
		},
	)
	if err != nil {
		panic(err)
	}

	C.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting",request.URL)
	})

	C.OnHTML("a[href]", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		fmt.Println(link)
		err := C.Visit(element.Request.AbsoluteURL(link))
		if err != nil{
			panic(err)
		}
	})

	C.OnResponse(func(response *colly.Response) {
		log.Println("response accpet ", response.StatusCode)
	})

	C.OnError(func(response *colly.Response, e error) {
		fmt.Println("Request URL: ", response.Request.URL,"err :", e)
	})

	for i := 0 ; i < 5 ; i++ {
		queue2.AddURL(fmt.Sprintf("%s?n=%d",url,i))
	}

	queue2.Run(C)
}
