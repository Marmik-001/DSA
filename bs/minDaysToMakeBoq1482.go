package bs

func enoughFlowers(nums []int , boq int , flower int , days int) bool {

	count := 0 


	for _, v := range nums {
		if v - days <= 0  {
			count++
			if count == flower {
				boq--
				if boq == 0  {
					return true
				}
				count = 0
			}
		} else {
			count = 0
		}
	}

	return false
}

func MinDaysToMake(bloomDay []int, m int, k int) int {

	ans := -1
	l := 1
	h := bloomDay[0]
	for _, v := range bloomDay {
		if v > h {
			h = v
		}
	}
	for l <= h {
		mid := (l + h) / 2
		if enoughFlowers(bloomDay , m , k , mid) {
			ans = mid
			h = mid - 1
		} else {
			l = mid + 1
		}

	}

	return ans
}