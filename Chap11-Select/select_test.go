package selec

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compare speed", func(t *testing.T) {
		slowServer := makeServer(2 * time.Millisecond)
		fastServer := makeServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Didn't expect an error:%v", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns error if time bigger than 10s", func(t *testing.T) {
		server := makeServer(12 * time.Millisecond)
		timer := 1 * time.Millisecond

		defer server.Close()

		_, err := ConfigRacer(server.URL, server.URL, timer)

		if err == nil {
			t.Errorf("Expected an error but got none")
		}
	})
}

func makeServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
