package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		Name     string
		Nums     []int
		Expected [][]int
	}{
		{
			Name: "Simple",
			Nums: []int{-1, 0, 1, 2, -1, -4},
			Expected: [][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := threeSum(test.Nums)
			assert.ElementsMatch(t, test.Expected, actual, actual)
		})
	}
}
