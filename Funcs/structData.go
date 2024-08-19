package Groupie_tracker

type JsonData struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Artist struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	DatesLocations map[string][]string `json:"datesLocations"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Date           []string            `json:"dates"`
	Location       interface{}         `json:"locations"`
}
type Location struct {
	Location interface{} `json:"locations"`
}
type Date struct {
	Date []string `json:"dates"`
}
type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Errors struct {
	Message     string
	Code        int
	Description string
}

type AllMessageErrors struct {
	NotFound                 string `json:"notfound"`
	BadRequest               string `json:"badrequest"`
	InternalError            string `json:"internalerror"`
	DescriptionNotFound      string
	DescriptionBadRequest    string
	DescriptionInternalError string
}