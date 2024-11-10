package mocking

import (
	"fmt"
	"io"
	"time"
)

// 該介面可以讓程序睡著
type Sleeper interface {
	Sleep()
}

// 該結構有實作 Sleeper 介面，並且存有一個時間間隔
type ConfigurableSleeper struct {
	// 決定打印完睡眠持續時間
	duration time.Duration

	// 決定睡眠的方式
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// 這是 Setter 方法，可以設置睡眠的方式
// 另外這裡 sleep 和 interface 的 Sleep 方法名稱一樣，如果這裡要設置 Getter 的話，
// ConfigurableSleeper 中的 func(time.Duration) 就必須改名稱，不然會衝突
func (c *ConfigurableSleeper) SetSleep(sleep func(time.Duration)) {
	c.sleep = sleep
}

func (c *ConfigurableSleeper) SetDuration(duration time.Duration) {
	c.duration = duration
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
