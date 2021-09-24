package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
)
var visited = map[string]bool{}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.jinyongwang.net" ),
		colly.MaxDepth(1),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// 已访问的页面不需要再看一次
		if visited[link] {
			return
		}
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
		visited[link] = true
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
		f,err := os.Create("./datacollection"+r.Request.URL.Path)
		defer f.Close()
		if err !=nil {
			fmt.Println(err.Error())
		} else {
			_,err=f.Write(r.Body)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})

	c.Visit("https://www.jinyongwang.net/data/renwu/")
}