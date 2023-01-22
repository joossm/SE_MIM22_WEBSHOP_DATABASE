package main

import (
	"SE_MIM22_WEBSHOP_DATABASE/handler"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

func main() {
	var serveMux = http.NewServeMux()
	serveMux.HandleFunc("/init", handler.InitDatabase)
	printStartUP()
	handler := cors.Default().Handler(serveMux)
	server := &http.Server{
		Addr:              ":8450",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           handler,
	}
	log.Fatal(server.ListenAndServe())
}
func printStartUP() {
	log.Printf("\n\n\t" +
		"DATABASESERVICE" +
		"\n\n" +
		"About to listen on Port: 8450." +
		"\n" +
		"\nSUPPORTED REQUESTS:" +
		"\nGET:" +
		"Go to http://127.0.0.1:8450/init to initialise the Database.")
}
