package array

func RemoveDuplicatesMapSolution(nums []int) int {
	uniques := []int{}
	hashMap := make(map[int]bool)
	for i, v := range nums {
		if _, ok := hashMap[nums[i]]; ok {
			continue
		} else {
			hashMap[v] = true
			uniques = append(uniques, v)

		}
	}
	copy(nums, uniques)

	return len(hashMap)
}

func RemoveDuplicatesArraySolution(nums []int) int {

	i := 0
	for j:= 1; i < len(nums)-1 && j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1

}
