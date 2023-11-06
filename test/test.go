package hsd

import "testing"

func TestStringDistance(t *testing.T) {
	git := StringDistance("foo", "foh")
	want := 1
	if got != want {
		t.Fatalf("expected: %v, git %v", want, got)
	}
}
