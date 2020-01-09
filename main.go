package main

import (
	"fmt"
	"flag"
	"os"
	"os/signal"
	"context"
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
)
func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome home!")
}

func CreateStorageHandler(w http.ResponseWriter, r *http.Request){
}

func UpdateStorageHandler(w http.ResponseWriter, r *http.Request){
}

func DeleteStorageHandler(w http.ResponseWriter, r *http.Request){
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w,r)
	})
}

type authenticationMiddleware struct {
	tokenUsers map[string]string
}

func main() {
	if(len(os.Args) < 3){
		log.Printf("Usage: writar <address> <port>")
		return
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connectings to finish - e.g. 15s or 1m")

	log.Printf("Starting writar.\n")
	address := os.Args[1]
	port := os.Args[2]

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/storage/", CreateStorageHandler).Methods("POST")
	router.HandleFunc("/storage/{id}", UpdateStorageHandler).Methods("PUT")
	router.HandleFunc("/storage/{id}", DeleteStorageHandler).Methods("DELETE")
	router.Use(loggingMiddleware);

	server := &http.Server{
		Handler: router,
		Addr: address+":"+port,
		WriteTimeout: 15 *  time.Second,
		ReadTimeout: 15 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	log.Printf("Start serving at " + address + " port: " + port)

	go func(){
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}
