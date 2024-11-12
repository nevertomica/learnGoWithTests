package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

// 這裡要測試的是 CheckWebsites 的行為，看他有沒有呼叫 mockWebsiteChecker
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
