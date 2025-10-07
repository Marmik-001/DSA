package slidingwindow

import "fmt"

func MaxPointsInCards(arr []int, k int) int {

	size := len(arr)
	high:= size - 1
	low := k -1
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	currentSum := sum
	fmt.Println("initial sum : " ,sum)
	for i := 0; i < k; i++ {
		currentSum = currentSum + arr[high] - arr[low]
		fmt.Println("current sum: ",currentSum)
		high--
		low--
		sum = max(currentSum, sum)

	}

	return sum
}