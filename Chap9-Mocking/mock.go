package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	final     string = "Go!"
	countdown int    = 3
	sleep            = "sleep"
	write            = "write"
)

type Sleeper interface {
	Sleep()
}

/*
First iteration of the test sleeper

	type SpySleeper struct {
		Calls int
	}

	func (s *SpySleeper) Sleep() {
		s.Calls++
	}
First iteration of the app sleeper

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
*/

type SpyCountOps struct {
	Calls []string
}

func (s *SpyCountOps) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountOps) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleep Sleeper) {
	for i := countdown; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleep.Sleep()
	}
	fmt.Fprint(out, final)
}

func main() {
	sleep := &ConfigSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleep)
}
