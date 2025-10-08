package bs

func findPeakElement(nums []int) int {
    n := len(nums)
    l, h := 0 , n - 1
    mid := (l + h)/2
    if n == 1 {
        return 0
    }
    for l <= h {
        mid = (l + h)/2
        if mid == 0 || mid == n - 1 {
			if mid == 0 && nums[mid] > nums[mid+1] {
                return mid
            } else if mid == n - 1 && nums[mid] > nums[mid - 1] {
                return mid
            } else if mid == 0 && nums[mid] < nums[mid+1] {
				l = mid + 1
            } else {
                h = mid - 1
            }
        } else {
            if nums[mid]  > nums[mid + 1] && nums[mid] > nums[mid - 1] {
                return mid
            } else if nums[mid] < nums[mid + 1] {
                l = mid + 1
            } else {
                h = mid - 1
            }
        }      
    }
    return mid
}