package mocking

import (
	"bytes"
	"slices"
	"testing"
)

const sleep = "sleep"
const write = "write"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpySleeper struct {
	Calls int
}

func (sleeper *SpySleeper) Sleep() {
	sleeper.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("print 3", func(t *testing.T) {
		buffer := bytes.Buffer{}
		mockSleeper := SpySleeper{}

		Countdown(&buffer, &mockSleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if mockSleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", mockSleeper.Calls)
		}

	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := SpyCountdownOperations{}

		Countdown(&spySleepPrinter, &spySleepPrinter)

		got := spySleepPrinter.Calls
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, got) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}

	})
}
