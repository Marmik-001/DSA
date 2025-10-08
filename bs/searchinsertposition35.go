package bs

func SearchInsert(nums []int, target int) int {
    l , r := 0 , len(nums) - 1
    mid := (r+l)/2
    for l <= r {
        mid = (r+l)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }      

   return l
}