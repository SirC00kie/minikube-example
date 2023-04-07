package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello server1! \n %s", time.Now())
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello user!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/user/", userHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
