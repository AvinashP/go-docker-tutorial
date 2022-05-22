package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Go sample page!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "App has successfully been deployed to digital ocean.")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
