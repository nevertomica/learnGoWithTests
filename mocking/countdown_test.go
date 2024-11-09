package mocking

import (
	"bytes"
	"testing"
)

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
}
