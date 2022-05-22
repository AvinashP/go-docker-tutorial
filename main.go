package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go docker tutorial")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hellow world!")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
