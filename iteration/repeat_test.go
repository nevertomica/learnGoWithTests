package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeated input times less than 0", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := ""

		if repeated != expected {
			assertCorrectRepeat(t, expected, repeated)
		}

	})

	t.Run("specified repeated times", func(t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"

		if repeated != expected {
			assertCorrectRepeat(t, expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("c", 1000)
	}
}

func assertCorrectRepeat(t testing.TB, expected, real string) {
	t.Helper()
	if expected != real {
		t.Errorf("expected %q but got %q", expected, real)
	}

}

func ExampleRepeat() {
	repeat := Repeat("f", 7)
	fmt.Println(repeat)
	//Output: fffffff
}
