package main

import (
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"github.com/gorilla/mux"
	"log"
	_ "log"
	_ "math/rand"
	"net/http"
	_ "net/http"
	_ "strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", GetMovies).Methods("GET")
	router.HandleFunc("/movies", CreateMovies).Methods("POST")
	router.HandleFunc("/movies/{id}", UpdateMovies).Methods("PUT")
	router.HandleFunc("/movies/{id}", DeleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
