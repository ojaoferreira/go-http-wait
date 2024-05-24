package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/timeout/second/{second}", func(w http.ResponseWriter, r *http.Request) {
		param := r.PathValue("second")

		timeout, _ := strconv.Atoi(param)

		time.Sleep(time.Duration(timeout) * time.Second)

		json.NewEncoder(w).Encode(map[string]string{"message": "we got it after " + param + " seconds"})
	})

	http.HandleFunc("/server/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "Ok"})
	})

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
