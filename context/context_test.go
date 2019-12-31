package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		data := "hello, world"
		svr := Server(&SpyStore{data, false})

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, res.Body.String(), data)
		}
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &StubStore{response: data}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellintCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellintCtx)

		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
}
