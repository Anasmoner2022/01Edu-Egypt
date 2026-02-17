package main

import (
	"encoding/json"
	"fmt"
	"groupi_tracker/models"
	"log"
	"net/http"
)

func main() {
	artistData := `[{
    "id": 1,
    "image": "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
    "name": "Queen",
    "members": [
      "Freddie Mercury",
      "Brian May",
      "John Daecon",
      "Roger Meddows-Taylor",
      "Mike Grose",
      "Barry Mitchell",
      "Doug Fogie"
    ],
    "creationDate": 1970,
    "firstAlbum": "14-12-1973",
  }]`
	var artists []models.Artist
	err := json.Unmarshal([]byte(artistData), &artists)
	if err != nil {
		log.Fatalf("Error when extract json: %v", err)
	}
	fmt.Println("=== Artist Struct Verification ===")
	// %+v prints field names + values
	fmt.Printf("%#v\n\n", artists[0])
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
