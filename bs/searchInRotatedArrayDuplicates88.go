package bs

func SearchInRotateArrDuplicates(nums []int, target int) bool {
	n := len(nums)

	l , h := 0  , n - 1
	for l <= h  {
		mid := (l + h) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] == nums[l] && nums[mid] == nums[h] {
			l++
			h--
		} else if nums[mid] < nums[l] {
			if target <= nums[h] && target > nums[mid] {
				l = mid + 1
			} else {
				h = mid - 1
			}
		} else {
			if target >= nums[l] && target <= nums[mid] {
				h = mid - 1
			} else {
				l = mid + 1
			}
		}
	} 

	return false
}