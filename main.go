package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/idqam/urlShortenerAPI/shortener"
)

func main(){

	//new instance of mux router 
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/home", WriteHandler).Methods("GET") // to make sure it only allows get request to this endpoint 
	fmt.Println("Server running on port 8080" )
	log.Fatal(http.ListenAndServe(":8080", r))
	

}

func WriteHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("In the response"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//obj := shortener.GenerateShortCode(6,"https://gobyexample.com/structs")
	
	
	fmt.Fprintln(w, 2+2)
	fmt.Fprintln(w, shortener.ValidateURL("https://gobyexample.com/structs") )
}