package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

func generateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	return randomBytes, err
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/32", http.StatusMovedPermanently)
	}
	length, error := strconv.Atoi(r.URL.Path[1:])
	if error != nil {
		handleError(w, error)
		return
	}
	if length <= 0 {
		handleError(w, errors.New("length must be greater than 0"))
		return
	}
	randomBytes, err := generateRandomBytes(length / 2)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Write([]byte(hex.EncodeToString(randomBytes)))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	http.HandleFunc("/", indexHandler)
	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
