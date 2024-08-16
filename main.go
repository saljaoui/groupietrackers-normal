package main

import (
	"fmt"
	"log"
	"net/http"

	Groupie_tracker "groupie_tracker/Funcs"
)

func main() {
	http.HandleFunc("/", Groupie_tracker.GetDataFromJson)
	http.HandleFunc("/Artist/{id}", Groupie_tracker.HandlerShowRelation)
	http.HandleFunc("/styles/", Groupie_tracker.HandleStyle)
	fmt.Println("http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
