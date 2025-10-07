package slidingwindow

func MaxConsiqutiveOnes(arr []int , k int) int {

	maxLength := 0
	zeroes := 0
	n := len(arr)
	l := 0
	
	for r:=0; r < n ; r++ {
		if arr[r] == 0 {
			zeroes++
		
			if zeroes > k {
				for zeroes  > k  {
					if arr[l] == 0 {
						zeroes--
					} 
					l++
				}
			}
		}
		
		maxLength = max(maxLength, r - l + 1)
		
	}


	return maxLength

}