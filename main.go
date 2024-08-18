package main

import (
	"fmt"
	"log"
	"net/http"

	Groupie_tracker "groupie_tracker/Funcs"
)

func main() {
	port := ":8085"
	http.HandleFunc("/", Groupie_tracker.GetDataFromJson)
	http.HandleFunc("/Artist/{id}", Groupie_tracker.HandlerShowRelation)
	http.HandleFunc("/styles/", Groupie_tracker.HandleStyle)
	fmt.Printf("http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
