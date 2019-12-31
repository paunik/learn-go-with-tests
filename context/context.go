package context

import "net/http"

import "fmt"

// Store interface
type Store interface {
	Fetch() string
}

// Server provides http.HandlerFunc
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(w, store.Fetch())
	}
}
