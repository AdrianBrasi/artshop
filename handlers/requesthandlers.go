package handlers

import (
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}

func ReturnHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the homepage</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the gallery. The art will live here</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the 'About' page</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the 'Contact' page</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<p>This is the 'Cart' page</p>"))
	if err != nil {
		log.Println("error writing response", err)
	}
}
