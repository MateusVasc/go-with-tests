package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Say 'a' 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})
}
