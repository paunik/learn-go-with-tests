package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store interface
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server provides http.HandlerFunc
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())
		fmt.Fprintf(w, data)
	}
}
