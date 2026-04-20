package bs

import (
	"fmt"
	"math"
	"slices"
)

func resFunc( nums []int , thresold int , mid int) bool {
	counter := 0
	for _ , v := range nums {
		counter += int(math.Ceil(float64(v) / (float64(mid))))
	}
	return counter <= thresold
}

func FindSmallestDivisorGreaterOrEqualToThresold(nums []int, thresold int) int { 

	if len(nums) > thresold {
		return -1
	}
	h := slices.Max(nums)
	l := 1
	ans := 0
	for l <= h {
		mid := ( l + h ) / 2
		res := resFunc(nums,  thresold  , mid)
		if res {
			fmt.Println("mid , ans : " , mid , ans)
			ans = mid
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}