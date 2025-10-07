package recursion
var ans [][]int
// func recursionDice(dice int, target int , k int) int {

// 	if target == 0 {
// 		return 0
// 	}
// 	if target < 0 {
// 		return -1
// 	}

// 	for face := 1 ; face < k ; face++ {
		
// 	}

// 	return 0
// }
func DiceProblem(dices int, sides int, target int) int{

	if dices == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	result := 0
	for face := 1; face <= sides; face++ {
		result += DiceProblem(dices - 1, sides, target - face)
	}
	return result 
}