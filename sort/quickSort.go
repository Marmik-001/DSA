package sort

// import "fmt"

func sort(arr []int , low int, high int) int {

	pivot := arr[low]
	l := low 
	r := high
	for l < r {
		
		for arr[l] <= pivot && l < high {
			l++
		}
		for arr[r] > pivot && r >= low {
			r--
		}
		if l < r {
			arr[l] , arr[r] = arr[r] , arr[l]
		}
	}

	arr[low] , arr[r] = arr[r] , arr[low]

	return r

}


func QuickSort(arr []int, low int, high int) {
	if low < high {
		partition := sort(arr, low , high)

		QuickSort(arr , low , partition - 1)
		QuickSort(arr , partition + 1 ,high)
	}
}