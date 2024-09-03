package selec

import (
	"fmt"
	"net/http"
	"time"
)

const timer = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigRacer(a, b, timer)
}

func ConfigRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

/* sync way
func measureTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return duration
}
*/
