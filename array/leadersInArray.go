package array

import "slices"

func LeadersInArray(nums []int) []int {
	//your code goes here
	n := len(nums)
	currentLeader := n - 1
	arr := []int{}
	arr = append(arr, nums[currentLeader])
	for i := n - 2; i > 0; i-- {
		if nums[i] > nums[currentLeader] {
			arr = append(arr, nums[i])
			currentLeader = i
		}
	}

	slices.Reverse(arr)

	return arr
}