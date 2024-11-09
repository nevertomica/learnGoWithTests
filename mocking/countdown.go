package mocking

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

const countDownStart = 3
const lastWord = "Go!"

func Countdown(writer io.Writer, sleeper Sleeper) {

	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, lastWord)
}
