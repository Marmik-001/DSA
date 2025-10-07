package slidingwindow

func FruitsInBasket(fruits []int) int {

	maxFruits := 0
	l := 0
	n := len(fruits)
	hashMap := make(map[int]int)
	
	for r := 0; r < n; r++ {
		hashMap[fruits[r]]++

		if len(hashMap) > 2 {
			hashMap[fruits[l]]--
			if hashMap[fruits[l]] == 0 {
				delete(hashMap , fruits[l])
			}
			l++
		}

		
		maxFruits = max(maxFruits , r - l + 1)
		
		
	}


	return maxFruits

}




// for r := 0; r < n; r++ {

		

// 		if _ , ok := hashMap[fruits[r]]; ok {
// 			if len(hashMap) > 2 {
// 				hashMap[fruits[l]]--
// 				if hashMap[fruits[l]] == 0 {
// 					delete(hashMap , fruits[l])
// 				}
// 				hashMap[fruits[r]]++
// 				l++
// 			} else {
// 			hashMap[fruits[r]]++
// 			}
// 		} else {
// 			if len(hashMap) < 2 {
// 				hashMap[fruits[r]]++
// 			} else {
// 				hashMap[fruits[l]]--
// 				if hashMap[fruits[l]] == 0 {
// 					delete(hashMap , fruits[l])
// 				}
// 				hashMap[fruits[r]]++
// 				l++
// 			}
// 		} 
		



// 		if len(hashMap) <= 2 {
// 			maxFruits = max(maxFruits , r - l + 1)
// 		}
		
// 	}