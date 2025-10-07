package slidingwindow

func SubstringWithKDistinctCharacter(s string , k int) int {
	maxString := 0
	l := 0
	n := len(s)
	hashMap := make(map[byte]int)
	
	for r := 0; r < n; r++ {
		hashMap[s[r]]++
		

		if len(hashMap) > k {
			hashMap[s[l]]--
			if hashMap[s[l]] == 0 {
				delete(hashMap , s[l])
			}
			l++
		}

		
		maxString = max(maxString , r - l + 1)
		
		
	}


	return maxString
}