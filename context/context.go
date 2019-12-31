package context

import "net/http"

import "fmt"

// Store interface
type Store interface {
	Fetch() string
	Cancel()
}

// Server provides http.HandlerFunc
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprintf(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
