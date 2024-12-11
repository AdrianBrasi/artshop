package main

import (
	"artshop/handlers"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/home", handlers.ReturnHomeHandler)
	http.HandleFunc("/gallery", handlers.GalleryHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	http.HandleFunc("/cart", handlers.CartHandler)

	log.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Artwork struct {
	ID          int
	Title       string
	Description string
	ImageURL    string
	Price       float64
	Available   bool
}
