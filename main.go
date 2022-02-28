package main

import (
	"fmt"
	"ramazan/helpers"
	"sync"

	"github.com/gocolly/colly"
)

//here im parsing 3 pages of artists and putting results in json format
func main() {
	wg := sync.WaitGroup{}
	file := helpers.OpenFile("listOfArtist.json")
	defer file.Close()

	sliceOfArtist := []string{}
	sliceOfBio := []string{}

	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error occured on this url %s \nerror:%s", r.Request.URL, e)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("We are visiting this site:", r.Request.URL)
	})

	c.OnHTML("h3.big-artist-list-title>a.link-block-target", func(h *colly.HTMLElement) {
		sliceOfArtist = append(sliceOfArtist, h.Text)
	})

	c.OnHTML(".big-artist-list-bio>p:first-child", func(h *colly.HTMLElement) {
		sliceOfBio = append(sliceOfBio, h.Text)
	})

	c.Visit("https://www.last.fm/ru/tag/pop/artists")

	for i := 2; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Visit(fmt.Sprintf("https://www.last.fm/ru/tag/pop/artists?page=%d", i))
		}(i)
	}
	wg.Wait()

	helpers.WriteToFile(sliceOfArtist, sliceOfBio, file.Name())
}
