package strutil

import (
	"testing"
)

func TestReversing(t *testing.T) {
	str := "abcdefg"
	if Reverse(str) != "gfedcba" {
		t.Errorf("reverse(%s) = %s, want %s", str, Reverse(str), "gfedcba")
	}
}
