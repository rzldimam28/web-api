package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is running...")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	hello := []byte("API w Ubuntu")
	w.Write(hello)
}