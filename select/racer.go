package selects

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// 測試 a, b 兩個 url 的回應時間，回傳較快的 url
func Racer(a, b string) (winnter string, err error) {
	return ConfigureRacer(a, b, tenSecondTimeout)
}

// 先有 Racer 測通 happy path，後面才需要測試超時的 sad path
// 而 happy path 不管會不會超時，又希望在測試時候可以最短時間測試超時
// 設計出 ConfigureRacer 來讓測試時，可以調整極短時限
// 這樣就可以不改動 Racer 的情況下，測試超時的 sad path，
// 其中 happy path 預設時限 10 秒
func ConfigureRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// 透過 select channel 的操作，和官方 time.After、time.Tick 一樣
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)

	}
}

// 在有 Http Response 之後，關閉 channel
// 透過關閉 channel 時候的回傳值，來判段哪一個 url 最先回應
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// 測試 url 的回應時間
// func measureResponseTime(url string) time.Duration {
// 	startAt := time.Now()
// 	http.Get(url)
// 	return time.Since(startAt)
// }
