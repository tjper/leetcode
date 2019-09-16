package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		Name         string
		Nums         []int
		Target       int
		ExpectedPair []int
	}{
		{
			Name:         "Simple",
			Nums:         []int{2, 7, 11, 15},
			Target:       9,
			ExpectedPair: []int{0, 1},
		},
		{
			Name:         "Duplicates",
			Nums:         []int{3, 3},
			Target:       6,
			ExpectedPair: []int{0, 1},
		},
		{
			Name:         "Complex",
			Nums:         []int{1, 4, 12, 14, 9, 6, 6},
			Target:       12,
			ExpectedPair: []int{5, 6},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := twoSum(test.Nums, test.Target)
			assert.ElementsMatch(t, test.ExpectedPair, actual)
		})
	}
}

func benchmarkTwoSum(b *testing.B, nums []int, target int) {
	for n := 0; n < b.N; n++ {
		twoSum(nums, target)
	}
}
func BenchmarkTwoSum1(b *testing.B) { benchmarkTwoSum(b, []int{2, 7, 11, 15}, 9) }
func BenchmarkTwoSum2(b *testing.B) { benchmarkTwoSum(b, []int{3, 3}, 6) }
func BenchmarkTwoSum3(b *testing.B) { benchmarkTwoSum(b, []int{1, 4, 12, 14, 9, 6, 6}, 12) }

func TestTwoSumFaster(t *testing.T) {
	tests := []struct {
		Name         string
		Nums         []int
		Target       int
		ExpectedPair []int
	}{
		{
			Name:         "Simple",
			Nums:         []int{2, 7, 11, 15},
			Target:       9,
			ExpectedPair: []int{0, 1},
		},
		{
			Name:         "Duplicates",
			Nums:         []int{3, 3},
			Target:       6,
			ExpectedPair: []int{0, 1},
		},
		{
			Name:         "Complex",
			Nums:         []int{1, 4, 12, 14, 9, 6, 6},
			Target:       12,
			ExpectedPair: []int{5, 6},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := twoSumFaster(test.Nums, test.Target)
			assert.ElementsMatch(t, test.ExpectedPair, actual)
		})
	}
}

func benchmarkTwoSumFaster(b *testing.B, nums []int, target int) {
	for n := 0; n < b.N; n++ {
		twoSum(nums, target)
	}
}
func BenchmarkTwoSumFaster1(b *testing.B) { benchmarkTwoSumFaster(b, []int{2, 7, 11, 15}, 9) }
func BenchmarkTwoSumFaster2(b *testing.B) { benchmarkTwoSumFaster(b, []int{3, 3}, 6) }
func BenchmarkTwoSumFaster3(b *testing.B) { benchmarkTwoSumFaster(b, []int{1, 4, 12, 14, 9, 6, 6}, 12) }
