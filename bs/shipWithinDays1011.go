package bs

import (
	"slices"
)

func shipable(nums []int, days int, mid int) bool {

	currentLoad := 0
	reqDays := 1 
	for _ , weight := range nums {
		if weight > mid {
			return false
		}

		if currentLoad + weight <= mid {
			currentLoad += weight
		} else {
			currentLoad = weight 
			reqDays++
		}
	}

	return reqDays <= days
}

func ShipWithInDays(nums []int, days int) int {
	
	h := slices.Max(nums) * len(nums)
	l := 1
	ans := -1
	for l <= h {
		mid := (l + h) / 2

		if shipable(nums, days, mid) {
			ans = mid
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}
