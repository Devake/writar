package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome home!")
}

func main() {
	log.Printf("Starting writar.\n")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	address := "0.0.0.0"
	port := "8080"
	log.Printf("Start serving at " + address + " port: " + port)
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}
