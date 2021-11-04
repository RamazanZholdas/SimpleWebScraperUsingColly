package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

//Я не понимаю почему оно не записывает данных в json файл
func main() {
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

	for i := 1; i <= 5; i++ {
		c.Visit(fmt.Sprintf("https://www.last.fm/ru/tag/pop/artists?page=%d", i))
	}

	var artist Artist
	slice := []Artist{}
	for i := range sliceOfArtist {
		top := i + 1
		artist = Artist{
			numberInTop: top,
			name:        sliceOfArtist[i],
			bio:         sliceOfBio[i],
		}
		slice = append(slice, artist)
	}

	for i := range slice {
		file, err := json.MarshalIndent(slice[i], "", " ")
		if err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile("artists.json", file, 0644); err != nil {
			log.Fatal(err)
		}
	}

	for i := range slice {
		fmt.Println(slice[i])
	}
}

type Artist struct {
	numberInTop int
	name        string
	bio         string
}
