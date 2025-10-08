package bs

func FindMin(nums []int) int {
    n := len(nums)
    l , h := 0 , n - 1
    if n == 1 || nums[0] < nums[h] {
        return nums[0]
    }
    for l <= h {
        mid := (l + h) / 2
        if mid == n - 1 {
            return mid
        } else if nums[mid] > nums[mid + 1] {
            return nums[mid + 1]
        } else if nums[mid] < nums[mid + 1] && nums[mid] < nums[mid - 1] {
            return nums[mid]
        } else if nums[mid] > nums[0] && nums[mid] >  nums[n-1] {
            l = mid + 1
        } else {
            h = mid - 1
        }
    }
    return -1
}