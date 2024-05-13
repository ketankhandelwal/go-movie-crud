package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID        string    `json:"id"`
	Isbn      string    `json:"isbn"`
	Title     string    `json:"title"`
	Direcetor *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, value := range movies {
		if params["id"] == value.ID {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}

}

func getMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, items := range movies {
		if items.ID == params["id"] {
			json.NewEncoder(w).Encode(items)
			return

		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range movies {
		if items.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}

}

var movies []Movie

func main() {

	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "Kuan", Direcetor: &Director{
		FirstName: "Ketan", Lastname: "Khandelwal",
	}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getAllMovies).Method("GET")
	r.HandleFunc("/movie/{id}", getMovieByID).Method("GET")
	r.HandleFunc("/createMovie", createMovie).Method("POST")
	r.HandleFunc("/updateMovie/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/deleteMovies/{id}", deleteMovie).Method("DELETE")

	fmt.Println("Starting Server @ 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
