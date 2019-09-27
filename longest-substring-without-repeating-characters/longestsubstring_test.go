package longestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		Name     string
		Str      string
		Expected int
	}{
		{
			Name:     "1",
			Str:      "abcabcbb",
			Expected: 3,
		},
		{
			Name:     "2",
			Str:      "bbbbb",
			Expected: 1,
		},
		{
			Name:     "3",
			Str:      "pwwkew",
			Expected: 3,
		},
		{
			Name:     "4",
			Str:      " ",
			Expected: 1,
		},
		{
			Name:     "5",
			Str:      "au",
			Expected: 2,
		},
		{
			Name:     "6",
			Str:      "dvdf",
			Expected: 3,
		},
		{
			Name:     "7",
			Str:      "aabaab!bb",
			Expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			length := lengthOfLongestSubstring(test.Str)
			assert.Equal(t, test.Expected, length)
		})
	}
}
