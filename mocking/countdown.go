package mocking

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (this *SpySleeper) Sleep() {
	this.Calls++
}

const countDownStart = 3
const lastWord = "Go!"

func Countdown(writer io.Writer, sleeper Sleeper) {

	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
		//time.Sleep(1 * time.Second)
	}
	fmt.Fprint(writer, lastWord)
}
