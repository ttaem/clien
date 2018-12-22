package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/ttaem/clien/gui"
)

//Content ...
type Content struct {
	Symph  string
	Title  string
	Link   string
	Author string
	Hit    string
	Time   string
}

func main() {
	//contents := []Content
	contents := make([]Content, 0, 200)

	c := colly.NewCollector(
	//								colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	c.OnHTML("div[class=content_list]", func(e *colly.HTMLElement) {
		e.ForEach("div[data-role=list-row]", func(_ int, e2 *colly.HTMLElement) { /* class=symph_row is not working */
			fmt.Println("+++++++++++++++++++++++++++++++")
			fmt.Println(e2.ChildText("div[data-role=list-like-count]"))
			fmt.Println(e2.ChildText("span[data-role=list-title-text]"))
			//fmt.Println(e2.ChildText("div[class=list_author]"))
			fmt.Println(e2.ChildText("span[class=nickname]"))
			fmt.Println(e2.ChildText("span[class=hit]"))
			fmt.Println(e2.ChildText("span[class=timestamp]"))
			fmt.Println("===============================")
			content := Content{
				Symph:  e2.ChildText("div[data-role=list-like-count]"),
				Title:  e2.ChildText("span[data-role=list-title-text]"),
				Link:   e2.ChildAttr("a[class=list_subject]", "href"),
				Author: e2.ChildText("span[class=nickname]"),
				Hit:    e2.ChildText("span[class=hit]"),
				Time:   e2.ChildText("span[class=timestamp]"),
			}
			contents = append(contents, content)
		})

	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response", r.StatusCode, string(r.Body), "\nError", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	//c.Visit("https://www.clien.net/service/board/park")
	//fmt.Println(contents)

	g, err := gui.NewGui()
	if err != nil {
		log.Panicln(err)
		return
	}
	g.Run()
}
