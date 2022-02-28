package helpers

import (
	"encoding/json"
	"log"
	"os"
	"ramazan/structs"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func WriteToFile(sliceOfArtist, sliceOfBio []string, fileName string) {
	var artists structs.Artists

	for i := range sliceOfArtist {
		artists.Artist = append(artists.Artist, structs.Artist{
			Id:   i,
			Name: sliceOfArtist[i],
			Bio:  sliceOfBio[i],
		})
	}

	byteArr, err := json.MarshalIndent(artists, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(fileName, byteArr, 0644)
}
