package main

import "net/http"

var store = newURLStore()

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8080", nil)
}

// Redirect is
func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

// Add is
func Add(w http.ResponseWriter, r *http.Request) {

}

type urlStore struct {
	mapper map[string]string
}

func (u *urlStore) Get(key string) string {
	return "http://www.google.com"
}

// NewURLStore is
func newURLStore() *urlStore {
	store := new(urlStore)
	// store.Get
	return store
}
