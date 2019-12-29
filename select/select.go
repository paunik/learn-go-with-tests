package selectexample

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// Racer checks who returns first. Returns error if timeout of 10s
func Racer(urlOne, urlTwo string) (winner string, err error) {
	return ConfigurableRacer(urlOne, urlTwo, tenSecondTimeout)
}

// ConfigurableRacer checks who returns first. Returns error if timeout of 10s
func ConfigurableRacer(urlOne, urlTwo string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urlOne):
		return urlOne, nil
	case <-ping(urlTwo):
		return urlTwo, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlOne, urlTwo)
	}

}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
