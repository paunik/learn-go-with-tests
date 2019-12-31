package context

import "net/http"

import "fmt"

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(w, store.Fetch())
	}
}
