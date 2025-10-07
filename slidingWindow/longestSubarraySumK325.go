package slidingwindow

func SubarraySum(nums []int, k int) int {

	mx := 0
	i := 0
	currentSum := 0
	n := len(nums)
	for j := 0; j < n; j++ {
		currentSum = currentSum + nums[j]
		if currentSum > k {
			currentSum = currentSum - nums[i]
			i += 1
			if currentSum == k  {
				mx = max(j-i + 1, mx)
			}
		} else if currentSum == k {
			mx = max(j-i + 1, mx)
		} 

	}

	return mx
}

func SubarraySumPositiveBetter(nums []int,k int) int {
	maxLength := 0 
	mp := make(map[int]int)
	sum := 0 
	n := len(nums)
	// [1,2,3,1,1,1,5]
	

	for i := 0 ; i < n ; i++ {
		if sum == k {
			maxLength = max(maxLength , i + 1)
		}
		
		
		diff := sum - nums[i]
		if j , ok := mp[diff] ; ok {
			le := i - j + 1
			maxLength =  max(maxLength , le)
		}
		sum += nums[i]
		if _ , ok := mp[sum] ; !ok {
			mp[sum] = i
		}
	} 

	return maxLength

}