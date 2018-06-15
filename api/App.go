package main

import (
	"log"
	"net/http"
)

func main() {
	err := IntializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
