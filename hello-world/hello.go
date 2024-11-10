package main

import (
	"fmt"
	"hello/mocking"
	"io"
	"net/http"
	"os"
	"time"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a personalised greeting in a given language.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// 這是第一次看到 Golang 的 http 處理
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

type DefaultSleeper struct {
}

func (defaultSleeper *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	// 這是第一次看到 Golang 的 log
	//log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))

	// sleeper := DefaultSleeper{}
	// mocking.Countdown(os.Stdout, sleeper)

	// 抽換成 ConfigurableSleeper
	sleeper := &mocking.ConfigurableSleeper{}
	sleeper.SetSleep(time.Sleep)
	sleeper.SetDuration(3 * time.Second)
	mocking.Countdown(os.Stdout, sleeper)

}
