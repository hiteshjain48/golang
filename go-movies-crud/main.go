package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"encoding/json"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies = []Movie{}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]... )
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {

		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			break
		}
	}
}
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn: "443233", Title: "Movie One",Director: &Director{FirstName: "John", LastName: "Doe"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("server on 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}
