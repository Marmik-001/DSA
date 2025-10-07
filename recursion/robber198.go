package recursion

func recursionRob(nums []int, index int) int {
    if(index >= len(nums)){
        return 0
    }

    // rob
    ansRob := nums[index] + recursionRob(nums, index + 2)

    // don't rob
    ansNotRob := 0 + recursionRob(nums, index + 1)

    return max(ansRob, ansNotRob)
}


func rob(nums []int) int {
    ans := recursionRob(nums, 0)
    return ans
}