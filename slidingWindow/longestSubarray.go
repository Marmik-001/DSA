package slidingwindow

func LongestSubarray(arr []int, k int) int {

	result := 0
	low := 0
	high := 0
	sum := 0
	for high < len(arr) {
		sum += arr[high]

		if sum > k && high < len(arr)-1 {

			sum = sum - arr[low]
			low++
			high++

		}
		if sum <= k {
			result = max(result, high-low+1)
			high++
		} 
	}

	return result
}
