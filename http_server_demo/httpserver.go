package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Print(header)
	for k, v := range header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	w.WriteHeader(200)

	responseCode := 200
	host := r.Host

	log.Printf("Status Code", responseCode)
	log.Printf("host", host)

}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	fmt.Println("start")
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8082", mux)
	if err != nil {
		log.Fatal(err)
	}
}
