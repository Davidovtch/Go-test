package sync

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("count to 3", func(t *testing.T) {
		counter := Counter{}
		want := 3

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, want)
	})

	t.Run("concurrently", func(t *testing.T) {
		wantedCount := 1000
		/*by using this constructor you don't need to
		  pass the memory address on the assertion*/
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, wantedCount)
	})
}
