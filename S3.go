package main
//Scapter Server ||| URL filter
import (
	"fmt"
	"github.com/go-acme/lego/log"
	"github.com/gocolly/colly"
	"net/http"
	"regexp"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	Url := r.URL.Query().Get("url")
	if Url == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting",Url)
	C := colly.NewCollector(
		//examples ---- URL filter
		colly.URLFilters(
			regexp.MustCompile("http://httpbin\\.org/(|e.+)$"),
			regexp.MustCompile("http://httpbin\\.org/h.+"),
		),
	)

	C.OnHTML("a[href]", func(element *colly.HTMLElement) {
		link := element.Request.AbsoluteURL(element.Attr("href"))
		fmt.Println(link)
	})

	C.OnResponse(func(response *colly.Response) {
		log.Println("response received", response.StatusCode)
	})

	err := C.Visit(Url)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write([]byte{123})
}


func main() {
	address := ":8888"
	http.HandleFunc("/",handler)
	log.Println("Listening on 8888")
	log.Fatal(http.ListenAndServe(address,nil))
}