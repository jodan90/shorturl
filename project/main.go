package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

var urlStore = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

func generateShortURL() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()

	urlStore.Lock()
	urlStore.m[shortURL] = longURL
	fmt.Printf("Stored: %s -> %s\n", shortURL, longURL)
	urlStore.Unlock()

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	tmpl.Execute(w, struct{ ShortURL string }{ShortURL: "http://localhost:8080/s/" + shortURL})
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/s/"):]
	fmt.Printf("Requested short URL: %s\n", shortURL)

	urlStore.RLock()
	longURL, ok := urlStore.m[shortURL]
	urlStore.RUnlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	fmt.Printf("Redirecting to: %s\n", longURL)
	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/shorten", shortenURLHandler)
	http.HandleFunc("/s/", redirectHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	fmt.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
