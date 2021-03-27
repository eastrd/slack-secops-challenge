package main

import (
	"log"
	"net/http"
	"os"
)

var (
	// BASICUSER : Basic Auth Username
	BASICUSER string
	// BASICPASS : Basic Auth Password
	BASICPASS string
	// CERTFILE : Path to cert file
	CERTFILE string
	// KEYFILE : Path to key file
	KEYFILE string
)

func init() {
	// Check environment variables are given
	BASICUSER = os.Getenv("BASIC_USER")
	BASICPASS = os.Getenv("BASIC_PASS")
	CERTFILE = os.Getenv("CERT")
	KEYFILE = os.Getenv("KEY")

	if len(BASICUSER) == 0 || len(BASICPASS) == 0 {
		log.Panic("Basic Username/Password is empty")
	}
	if len(CERTFILE) == 0 || len(KEYFILE) == 0 {
		log.Panic("Cert/Key is empty")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getwordfreq", handleGetWordFrequency)

	log.Println("server started at port 443")

	err := http.ListenAndServeTLS(":443", CERTFILE, KEYFILE, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check Basic Auth before /getwordfreq
		u, p, ok := r.BasicAuth()
		if !ok {
			log.Println("error parsing basic auth")
			w.WriteHeader(401)
			return
		}
		if u != BASICUSER || p != BASICPASS {
			log.Println("basic auth login failed")
			w.WriteHeader(401)
			return
		}
		mux.ServeHTTP(w, r)
	}))

	if err != nil {
		log.Fatalf("error starting server %s", err.Error())
	}
}
