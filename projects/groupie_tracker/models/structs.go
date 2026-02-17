package models

// Artist represents a band/artist from the /api/artists endpoint
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// Locations represents an individual location entry
type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Dates represents an individual date entry
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Relation links locations and dates via a map
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Index holds the wrapper for the locations, dates, and relations endpoints
// These endpoints return an object with an "index" key containing an array
type Index struct {
	Locations []Locations `json:"index"`
	Dates     []Dates     `json:"index"`
	Relation  []Relation  `json:"index"`
}
