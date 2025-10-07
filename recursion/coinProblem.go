package recursion

import (
	"fmt"
	"math"
)

func recursion(arr []int, target int) int {

	if target == 0 {
		return 0
	}
	if target < 0 {
		return math.MaxInt
	}
	minCoins := math.MaxInt
	fmt.Println(minCoins)
	for _, coin := range arr {
		fmt.Println("coin", coin)
		res := recursion(arr, target-coin)
		if res != math.MaxInt {
			minCoins = min(minCoins, res+1)
		}
	}
	fmt.Println(minCoins)
	return minCoins

}

func CoinProblemRecursion(coins []int, target int) {
	ans := recursion(coins, target)

	if ans == math.MaxInt {
		fmt.Println("Not possible")
	} else {
		fmt.Println("Minimum coins:", ans) // should be 3 (5+5+1)
	}
}


func recursionAns(coins []int, amount int) int {
    if amount == 0{
        return 0
    }
    ans := math.MaxInt
    minCoins := math.MaxInt

    for _, coin := range coins {
        if coin <= amount {
            ans =  recursion(coins, amount - coin)
            if ans != math.MaxInt {
                ans = ans + 1
            }
            minCoins =  min(ans , minCoins)
        }
    }
    return minCoins

}

func coinChange(coins []int, amount int) int {
    finalAns := recursionAns(coins, amount)
    if(finalAns == math.MaxInt){
        return -1
    }else {
        return finalAns
    }
}