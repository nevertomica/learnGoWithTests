package mocking

import (
	"bytes"
	"slices"
	"testing"
	"time"
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

// interface guard 可以檢查是否有實作特定 interface，此處是 Sleeper
// var _ Sleeper = (*SpyTime)(nil)

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.sleep}
	sleeper.Sleep()

	// 檢查 ConfigurableSleeper 的 Sleep 有無調用，如果有 SpyTime 則會設置 time Duration
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
