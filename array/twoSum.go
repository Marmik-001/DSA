package array

import "fmt"
func TwoSum(nums []int, target int) []int {
    
	fmt.Println("welcome to 2 sum solution")

    set := make(map[int]int)
    for i,v := range nums{
        diff := target - v
        if j, ok := set[diff];ok {
            return []int{i,j}
        }
        set[v] = i
    }
    return nil

}