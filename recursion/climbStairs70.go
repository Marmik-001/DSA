package recursion

var valueMap = make(map[int]int)

func ClimbStairs(n int) int {
    if n == 1 {
        return 1
    }else if n == 2 {
        return 2
    }else if val, ok := valueMap[n]; ok {
        return val
    }

    ans := ClimbStairs(n-1) + ClimbStairs(n-2)
    valueMap[n] = ans
    return ans
}