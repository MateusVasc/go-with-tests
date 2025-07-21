package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Adding to positive numbers", func(t *testing.T) {
		got := add(5, 5)
		want := 10

		if got != want {
			t.Errorf("Got %d want %d", got, want)
		}
	})
}
