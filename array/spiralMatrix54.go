package array

func SpiralOrder(m [][]int) []int {
	ans := []int{}
	c, r := len(m), len(m[0])
	min := min(r, c)
	spirals := 0
	if min%2 == 0 {
		spirals = min / 2
	} else {
		spirals = (min + 1) / 2
	}

	current := 0
	// i , j := 0, 0
	for current < spirals {
		for i, j := current, current; j < r-current; j++ {
			ans = append(ans, m[i][j])
		}

		if current+1 < c-current {

			for i, j := current+1, r-current-1; i < c-current; i++ {
				ans = append(ans, m[i][j])
			}
		} else {
			break
		}
		if r-current-2 > current-1 {

			for i, j := c-current-1, r-current-2; j > current-1; j-- {
				ans = append(ans, m[i][j])
			}
		}else {
			break
		}

		if c-current-2 > current {

			for i, j := c-2-current, current; i > current; i-- {
				ans = append(ans, m[i][j])
			}
		}else {
			break
		}
		current++
	}
	return ans
}
