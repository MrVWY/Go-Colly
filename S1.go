package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

//basic
func main() {
	//father
	C := colly.NewCollector(
		colly.AllowedDomains("hackerspaces.org"),
	)
	//childen
	C1 := C.Clone()

	C.OnRequest(func(request *colly.Request) {
		fmt.Println("Visit ", request.URL.String())
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
		response.Ctx.Put("Custom-header", response.Headers.Get("Custom-Header"))
		C1.Request("GET", "", nil, response.Ctx, nil)
		log.Println("response accpet ", response.StatusCode)
	})

	C.OnError(func(response *colly.Response, e error) {
		fmt.Println("Request URL: ", response.Request.URL,"err :", e)
	})

	_ = C.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "",
		Delay:        0,
		RandomDelay:  0,
		Parallelism:  2, //Limit the number of threads started by colly to two
	})

	_ = C.Visit("https://hackerspaces.org/")
}
