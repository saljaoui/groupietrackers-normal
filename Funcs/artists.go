package Groupie_tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// func to Get All data From json
func GetArtistsDataStruct() ([]JsonData, error) {
	var artistData []JsonData

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, fmt.Errorf("khata2 f jib l'data: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code machi normal: %d", response.StatusCode)
	}

	err = json.NewDecoder(response.Body).Decode(&artistData)
	if err != nil {
		return nil, fmt.Errorf("khata2 f t7wil JSON: %v", err)
	}
	return artistData, nil
}

// / func to fetching data from any struct and return Struct Artist with Id user
func FetchDataRelationFromId(id string) (Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api"
	urlartist := url + "/artists/" + id
	var artist Artist
	err := GetanyStruct(urlartist, &artist)
	if err != nil {
		return Artist{}, fmt.Errorf("error fetching data from artist data: %w", err)
	}

	if artist.ID == 0 {
		return Artist{}, fmt.Errorf("error fetching data from artist data: %w", err)
	}

	var date Date
	urldate := url + "/dates/" + id
	errdate := GetanyStruct(urldate, &date)
	if err != nil {
		return Artist{}, fmt.Errorf("error fetching data from artist data: %w", errdate)
	}

	var location Location
	urlLocation := url + "/locations/" + id
	errlocations := GetanyStruct(urlLocation, &location)
	if err != nil {
		return Artist{}, fmt.Errorf("error fetching data from locations data: %w", errlocations)
	}
	var relation Relation
	urlrelation := url + "/relation/" + id
	errrelation := GetanyStruct(urlrelation, &relation)
	if err != nil {
		return Artist{}, fmt.Errorf("error fetching data from locations data: %w", errrelation)
	}
	artist.Location = location.Location
	artist.Date = date.Date
	artist.DatesLocations = formatLocations(relation.DatesLocations)
	return artist, nil
}

// func to UnmarshalJSON from any struct with send url and any
// return error for has any error
func GetanyStruct[T any](url string, result *T) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data from URL: %w", err)
	}
	defer response.Body.Close()
	// Decode the JSON response into the provided struct
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return fmt.Errorf("error decoding JSON data: %w", err)
	}
	return nil
}

// func To Format String To remove '_' or '-' and Capitaliz text
func formatLocations(locations map[string][]string) map[string][]string {
	formatted := make(map[string][]string, len(locations))
	for location, dates := range locations {
		formattedLoc := strings.Title(strings.NewReplacer("-", " ", "_", " ").Replace(location))
		formatted[formattedLoc] = dates
	}
	return formatted
}