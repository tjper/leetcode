package main

type num struct {
	index    int
	val      int
	disabled bool
}

func newNum(index, val int) num {
	return num{
		index:    index,
		val:      val,
		disabled: false,
	}
}

func newNums(nums []int) []num {
	var numSet = make([]num, len(nums))
	for i, num := range nums {
		numSet[i] = newNum(i, num)
	}
	return numSet
}

func threeSum(nums []int) [][]int {
	// For each num in nums, there exists a complement that adds to 0. Find a
	// pair that sums to equal this complement. This pair, appended to num, is
	// the three-sum.
	var (
		numSet   = newNums(nums)
		triplets [][]int
	)
	for i, num := range numSet {
		numSet[i].disabled = true
		pair := twoSum(numSet, -1*num.val)
		numSet[i].disabled = false
		if pair == nil { // pair DNE
			continue
		}
		triplets = append(triplets, append(pair, num.val))
	}
	return triplets
}

// twoSum finds two ints from nums that sum to equal target.
func twoSum(nums []num, target int) []int {
	var (
		numExists = make(map[int]struct{}, len(nums))
		pair      []int
	)
	for _, num := range nums {
		if num.disabled {
			continue
		}
		complement := target - num.val
		if _, ok := numExists[complement]; ok {
			pair = []int{num.val, complement}
		}
		numExists[num.val] = struct{}{}
	}
	return pair
}
