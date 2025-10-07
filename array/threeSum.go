package array

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	var (
		ans  = [][]int{}
		n = len(nums) - 1
		seen = make(map[string]bool) // acts like a Set
	)

	fmt.Println("welcome to 3 sum solution")

	sort.Ints(nums)

	for pointer := 0; pointer < n -1; pointer++ {
		if pointer > 0 && nums[pointer] == nums[pointer-1] {
			continue
		}
		low,high := pointer + 1 , n
		 
		for low < high {
			sum := nums[high] + nums[low] +nums[pointer]
			if sum == 0 {
				triplet := []int{nums[high] , nums[low] ,nums[pointer]}
				key := fmt.Sprint(triplet)
				if !seen[key] {
					ans = append(ans, []int{nums[pointer] , nums[low] ,nums[high]})
					seen[key] = true
				}
				for low < high && nums[low] == nums[low+1] {
					low++
				}
				// skip duplicate highs
				for low < high && nums[high] == nums[high-1] {
					high--
				}
				low++
				high--
			} else if sum < 0{
				low++
			} else {
				high--
			}
			
		}
	}

	return ans
}
