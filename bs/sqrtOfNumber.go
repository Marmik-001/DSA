package bs

func Sqrt(n int) int {
	l ,  h  , ans := 0 , n/2 , -1
	
	for l <= h {
		mid := (l + h) / 2
		temp := mid * mid
		if temp == n {
			return mid
		} else if temp < n {
			ans = max(ans , mid)
			l = mid + 1
		} else {
			// ans = min(ans , mid)
			h = mid - 1
		}
	}

	return ans
}