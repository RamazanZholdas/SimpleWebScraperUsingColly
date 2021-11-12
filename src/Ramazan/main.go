package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

//here im parsing 3 pages of artists and putting results in json format
func main() {
	file, err := os.Create("listOfArtist.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var artists Artists
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
		c.Visit(fmt.Sprintf("https://www.last.fm/ru/tag/pop/artists?page=%d", i))
	}

	for i := range sliceOfArtist {
		artists.Artist = append(artists.Artist, Artist{
			Id:   i,
			Name: sliceOfArtist[i],
			Bio:  sliceOfBio[i],
		})
	}

	byteArr, err := json.MarshalIndent(artists, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(file.Name(), byteArr, 0644)
}

type Artists struct {
	Artist []Artist `json:"artists"`
}

type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
