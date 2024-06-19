package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay01(t *testing.T) {
	assert.Equal(t, 2000, GetIncreasedCount(), "should be equal")
}

// func TestReverseRunes(t *testing.T) {
// 	cases := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{"Hello, 世界", "界世 ,olleH"},
// 		{"", ""},
// 	}
// 	for _, c := range cases {
// 		got := ReverseRunes(c.in)
// 		if got != c.want {
// 			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
// 		}
// 	}
// }
