package main

// twoSum takes a set of nums and a target to reach. A subset of two nums is
// sought to add up to target. The two nums that satisfy this condition are
// returned in a slice.
func twoSum(nums []int, target int) []int {
	numexists := make(map[int]int, len(nums))
	for i, num := range nums {
		numexists[num] = i
	}

	var twosumpair = make([]int, 2)
	for i, num := range nums {
		needednum := target - num
		j, ok := numexists[needednum]
		if j == i { // skip current number
			continue
		}
		if ok { // found a pair
			twosumpair[0], twosumpair[1] = i, j
		}
	}

	return twosumpair
}

func twoSumFaster(nums []int, target int) []int {
	var (
		numExists = make(map[int]int, len(nums))
		pair      []int
	)
	for i, num := range nums {
		complement := target - num
		j, ok := numExists[complement]
		if ok {
			pair = []int{i, j}
			break
		}
		numExists[num] = i
	}
	return pair
}
