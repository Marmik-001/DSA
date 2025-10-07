package array

// import "fmt"


func MaxArea(height []int) int {
    
    left := 0
    right := len(height) -1
    result := 0
    for left < right{
        area := (right - left) * min(height[right], height[left])
        if area > result {
            result = area
        }
        if height[left] < height[right] {
            left++
        }else{
            right--
        }
    }
    return result

}