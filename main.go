package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *director `json:"director"`
}
type director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func (m *Movie) isempty() bool {
	return m.Title == ""
}
func main() {
	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "911",
		Title:    "Sare jaha se acha",
		Director: &director{Firstname: "satya", Lastname: "Nadela"},
	})

	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "333",
		Title:    "kala pani",
		Director: &director{Firstname: "rishi", Lastname: "sunak"},
	})
	movies = append(movies, Movie{
		ID:       "3",
		Isbn:     "874",
		Title:    "namaste london",
		Director: &director{Firstname: "mahesh", Lastname: "bhatt"},
	})

	movies = append(movies, Movie{
		ID:       "4",
		Isbn:     "43321",
		Title:    "hoowdy india",
		Director: &director{Firstname: "james", Lastname: "mccantyer"},
	})
	movies = append(movies, Movie{
		ID:       "5",
		Isbn:     "2644",
		Title:    "inception",
		Director: &director{Firstname: "christopher", Lastname: "nolan"},
	})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getonemovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4500", r))
}

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getonemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, value := range movies {
		if value.ID == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	json.NewEncoder(w).Encode("no Movie found with given id")
}

func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Create new Movie")
	}
	var newmovieentry Movie
	_ = json.NewDecoder(r.Body).Decode(&newmovieentry)
	if newmovieentry.isempty() {
		json.NewEncoder(w).Encode("your request sent is empyt! write some data")
	}
	rand.Seed(time.Now().UnixNano())
	newmovieentry.ID = strconv.Itoa(rand.Intn(100))
	movies = append(movies, newmovieentry)
	json.NewEncoder(w).Encode(newmovieentry)
	return
}

func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var update1 Movie
			_ = json.NewDecoder(r.Body).Decode(&update1)
			value.Title = update1.Title
			value.Director = update1.Director
			value.Isbn = update1.Isbn
			movies = append(movies, value)
			json.NewEncoder(w).Encode(value)
			return
		}
	}
}

func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
