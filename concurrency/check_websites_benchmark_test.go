package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	//defer timeElapsed("BenchmarkCheckWebsites", b.N, time.Now())

	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

// 下方函數可以打印 b.N 的值
// func timeElapsed(name string, loop int, current time.Time) {
// 	fmt.Printf("%s took %s\n with b.N = %d\n", name, time.Since(current), loop)
// }
