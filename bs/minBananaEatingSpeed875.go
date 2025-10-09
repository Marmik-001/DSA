package bs

import (
	"fmt"
	"math"
)

func canEat(nums []int, h int, speed int) bool {
	fmt.Println("canEat val : " , h , speed)
	n := len(nums)
	for i, v := range nums {
		if h == 0 {
			break
		}
		if v%speed == 0 {
			temp := v / speed
			h = h - temp
			if h < 0 {
				return false
			}
		} else {
			temp := v / speed
			h = h - temp - 1
			if h < 0 {
				return false
			}
		}
		if i == n - 1{
			return true
		}
	}

	return false
}

func MinEatingSpeed(piles []int, h int) int {

	low := 1
	high := piles[0]
	for _, v := range piles {
		if v < low {
			low = v
		}
		if v > high {
			high = v
		}
	}
	fmt.Println(low , high)
	ans := math.MaxInt
	for low <= high {
		mid := (low + high) / 2
		if canEat(piles, h, mid) {
			ans = min(mid , ans)
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return ans

}