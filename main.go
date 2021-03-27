package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/getwordfreq", handleGetWordFrequency)

	log.Println("server started at port 8080")
	// err := http.ListenAndServeTLS(":8090", "../certs/cert.pem", "../certs/key.pem", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
