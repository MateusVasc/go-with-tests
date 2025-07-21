package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Adding to positive numbers", func(t *testing.T) {
		got := Add(5, 5)
		want := 10

		if got != want {
			t.Errorf("Got %d want %d", got, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(3, 5)
	fmt.Println(sum)
	// Output: 8
}
