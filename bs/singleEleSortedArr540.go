package bs

func SingleNonDuplicate(nums []int) int {
    n := len(nums)
    l , h := 0 , n - 1
    if n == 1 {
        return nums[0]
    }
    for l <=h {
        mid := (l+h)/2
        
        if mid == n - 1 || mid == 0 {
            return nums[mid]
        } else if nums[mid] != nums[mid + 1] && nums[mid] != nums[mid -1 ] {
            return nums[mid]
        } else if mid & 1 == 0 {
            if nums[mid] == nums[mid + 1] {
                l = mid + 1
            } else {
                h = mid - 1
            }
        } else {
            if nums[mid] == nums[mid - 1] {
                l = mid + 1
            } else {
                h = mid - 1
            }
        }
    }
    return -1
}