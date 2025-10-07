package array

func RotateByKelementsToLeft(arr []int , k int) []int {
	n := len(arr) - 1
	k = k%n
	rotate(arr , 0 , k-1)
	rotate(arr , k , n  )
	rotate(arr , 0 , n )
	return arr
}

func rotate(arr []int, l int , r int) {

	for l <= r {
		arr[l] , arr[r] = arr[r] , arr[l]
		l++
		r--
	}

	
}


func RotateByKelementsToRight (arr []int , k int) []int {

	n := len(arr) - 1
	k = k % (n + 1)
	rotate(arr , n - k + 1, n  )
	rotate(arr , 0 , n - k)
	rotate(arr , 0 , n)

	return arr
}