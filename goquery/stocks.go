package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".reviews-wrap article .review-rhs").Each(func(i int, s *goquery.Selection) {
		band := s.Find("h3").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

func main() {

	doc, err := goquery.NewDocument("http://q.10jqka.com.cn/stock/gn/smx")
	if err != nil {
		panic(err)
	}

	// ret, err := doc.Find("#head").Html()
	price := doc.Find(".price").Text()
	pcent := doc.Find(".pcent").Text()

	println(price)
	println(pcent)

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		if i == 9 {
			return
		}

		s.Find("td").Each(func(i int, ss *goquery.Selection) {
			if i == 0 {
			} else if i == 1 || i == 3 {
				println(ss.Find("a").Text())
			} else {
				println(ss.Text())
			}
		})
	})

}
