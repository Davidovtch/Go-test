package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("prints 3 and go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpyCountOps{}

		Countdown(buffer, spy)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before print", func(t *testing.T) {
		spy := &SpyCountOps{}
		Countdown(spy, spy)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, spy.Calls) {
			t.Errorf("wanted calls %v got %v", want, spy.Calls)
		}
	})
}

func TestConfigSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
