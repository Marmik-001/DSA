package array

func NextPermutation(nums []int)  {
    n := len(nums) - 1
    i, j := n,  n
    for i > 0  {
        if nums[i] > nums[i-1] {
            break
        } 
        i--
    }
    if i == 0 {
        l , h := 0 , n
        for  l < h {
            nums[l] , nums[h] = nums[h] , nums[l]
            l++ 
            h--
        }
        return
    }
    
    swapper := i - 1
    for nums[j] <= nums[swapper] {
        j--
    }
    nums[swapper] , nums[j] = nums[j] , nums[swapper]
    low ,  high := i , n 
    for low < high {
        nums[low] , nums[high] = nums[high] , nums[low]
        low++
        high--
    }
}