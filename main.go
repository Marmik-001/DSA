package main

import (
	// "fmt"
	// ran "main/random"
	// "fmt"
	// array "main/array"
	// linkedlist "main/linkedList"
	// concepts "main/concepts"
	"fmt"
	"main/array"
	// recursion "main/recursion"
	// sliding "main/slidingWindow"
	// sort "main/sort"
	// "math"
)

func cycleValues(str []string, f func(string)) {

	for _, v := range str {
		f(v)
	}
}
// func recursion(arr []int, index int, size int ) [][]int {
// 	var result [][]int
// 	var currentSubSet []int
	
// 	var sub func(index int)
// 	sub = func(index int){
// 		if index == size {
// 			subsetCopy := make([]int, len(currentSubSet))
//             copy(subsetCopy, currentSubSet)
//             result = append(result, subsetCopy)
//             return
// 		}
		
// 		currentSubSet = append(currentSubSet, arr[index])
// 		sub(index + 1)

// 		currentSubSet  = currentSubSet[:len(currentSubSet)-1]
// 		sub(index + 1)


// 	}
// 	sub(index)
// 	return result
// }
func recursionStr(str string, index int, size int ) []string {
	var result []string
	var currentSubSet string
	
	var sub func(index int)
	sub = func(index int){
		if index == size {
			var subsetCopy string 
			subsetCopy = currentSubSet
			result = append(result, subsetCopy)
			return
		}
		
		currentSubSet = currentSubSet + string(str[index])
		sub(index + 1)

		currentSubSet = currentSubSet[:len(currentSubSet)-1]
		sub(index + 1)


	}
	sub(index)
	return result
}


type car struct {
	model string
	color string
	year  int
}

func newCar(model string, color string, year int) car {

	c := car{
		model: model,
		color: color,
		year:  year,
	}
	return c

}

func (c *car) carDemoReciever() {

	c.model = "asdf"
	c.color = "asdf"
	c.year = 5555

}

// func twoSum(nums []int, target int) []int {

// 	set := make(map[int]int)
// 	for i,v := range nums {
// 		diff := target - v
// 		if j, ok := set[diff]; ok{
// 			return []int{i,j}
// 		}
// 		set[v] = i

func isPowerOfFour(n int) bool {
	if n== 1{
		return true;
	}
	if n%4!=0 {
		return false;
	}
	
	return isPowerOfFour(n/4);
}


func main() {
	// fmt.Println(isPowerOfFour(1))

	// names := []string{"hello", "world", "bye"}
	// cycleValues(names, getLen)
	// hello()



	// fmt.Println(ran.Name)
	// name := "oldName"
	// pName := &name
	// updateName(pName)
	// fmt.Println(name)

	// carInfo := newCar("senna 750", "black", 2020)
	// fmt.Println(carInfo)

	// carInfo.carDemoReciever()

	// fmt.Println("car info after update : ",carInfo)

	// fmt.Println(twoSum( []int{1,2,3,4,5}, 4 ))

	// twoSum := array.TwoSum([]int{1,2,3,4,5}, 4)
	// fmt.Println(twoSum)

	// threeSum := array.ThreeSum([]int{0,0,0})
	// fmt.Println(threeSum)

	// linkedlist.Basic()

	// duplicate := linkedlist.FindDuplicate([]int{1,2,2,3,4,5,6}) // inputs 1 to n and one number repeating
	// fmt.Println(duplicate)

	// maxWater := array.MaxArea([]int{1,8,6,5,4,2,8,3,7})
	// fmt.Println(maxWater)

	// primes := array.CountPrime(10)
	// fmt.Println(primes)

	// arr := []int{2,1,5,4,3,0,0}
	// array.NextPermutation(arr)
	// fmt.Println(arr)



	// concepts.StructDemo()

	// ans := array.RemoveDuplicatesArraySolution([]int{1,1,1,1,3,3,3})
	// fmt.Println(ans)

	// ans := array.RotateByKelementsToRight([]int{1,2,3,4,5,6,7} , 1)
	// fmt.Println(ans)
	// arr := []int{}
	// fmt.Println(len(arr))
	// fmt.Println(recursion(50))

	// fmt.Println(array.LeadersInArray([]int{1,2,5,3,1,2}))

	// array.SetZeroes([][]int{[]int{1,1,1} , []int{1,0,1} , []int{1,1,1}})
	fmt.Println(array.SpiralOrder([][]int{{1,2,3,4} , {5,6,7,8} , {9,10,11,12}})) 
	// fmt.Println(recursion.ClimbStairs(2))
	// fmt.Println(recursion([]int{10, 2, 3, 4}, 0, 4))
	// fmt.Println(recursionStr("helo" , 0 , 4))
	// recursion.CoinProblemRecursion([]int{11,8,6} , 11)
	// fmt.Println(recursion.RemoveAllOccurenceOfASubstring("ababcddddababc" , "abc"))
	// fmt.Println(recursion.DiceProblem(2,6,7))


	// fmt.Println(sliding.LongestSubarray([]int{1,2,3,4,5,10,1,1,1,1,1,1,1,1,1,1} , 10))	
	// fmt.Println(sliding.MaxPointsInCards([]int{1,2,3,4,5,6,5,4,3,1} , 4))
	// fmt.Println(sliding.LengthOfLongestSubstring("asdfghjklasdfghjklasdfghjklasdfghjkl"))
	// fmt.Println(sliding.MaxConsiqutiveOnes([]int{1,1,1,0,0,0,0,1,1,1,0} , 2))
	// fmt.Println(sliding.FruitsInBasket([]int{3,3,3,1,2,1,1,2,3,3,4}))
	// fmt.Println(sliding.SubstringWithKDistinctCharacter("aaabcbcbcaadaa" , 3))
	// fmt.Println(sliding.NumberOfSubstringsWithAllThreeChars("bbacba"))
	// fmt.Println(sliding.CharacterReplacement("AABABBA" , 2))
	// fmt.Println(sliding.CharacterReplacementOptimal("AABABBA" , 2))
	// arr := []int{2,5,1,1,1,1,1,1,1,2,2,2,2,5}
	// sort.MergeSort(arr , 0 , len(arr) - 1)
	// sort.QuickSort(arr,  0 , len(arr) - 1)
	// fmt.Println(arr)
	// fmt.Println(sliding.SubarraySumPositiveBetter([]int{2,0,0,0,0,0,3} , 3))




}