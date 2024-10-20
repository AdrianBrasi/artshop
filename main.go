package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home", returnHomeHandler)
	http.HandleFunc("/gallery", galleryHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/cart", cartHandler)

	log.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Artwork struct {
	ID 			int
	Title 		string
	Description string
	ImageURL 	string
	Price 		float64
	Available 	bool
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}

func returnHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the homepage</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the gallery. The art will live here</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the 'About' page</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the 'Contact' page</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the 'Cart' page</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}
