package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestHandler(t *testing.T) {
	data := "hello, world"
	t.Run("happy path returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, res.Body.String(), data)
		}
	})
	// t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
	// 	store := &SpyStore{response: data, t: t}
	// 	svr := Server(store)

	// 	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// 	cancellintCtx, cancel := context.WithCancel(req.Context())
	// 	time.AfterFunc(5*time.Millisecond, cancel)
	// 	req = req.WithContext(cancellintCtx)

	// 	res := httptest.NewRecorder()

	// 	svr.ServeHTTP(res, req)

	// 	store.assertWasCancelled()
	// })
}
