package bs

func bs(nums []int , l int, r int,  target int) (int,  int) {
	if l > r {
		return -1,-1
	}
	mid := (l + r)/ 2
	
	if nums[mid] == target {
		finalLow , _ := bs(nums , l , mid - 1 , target )
		_ ,finalHigh := bs(nums , mid + 1,  r , target)
		if finalHigh == -1 {
			finalHigh = mid
		}
		if finalLow == -1 {
			finalLow = mid
		}
		return finalLow , finalHigh
	} else if  nums[mid] > target {
		return bs(nums , l ,mid- 1, target)
	} else {
		return bs(nums , mid + 1 , r , target)
	}
}

func SearchRange(nums []int, target int) []int {

    lb , ub := bs(nums, 0 , len(nums) - 1 , target)
    return []int{lb,ub}
}