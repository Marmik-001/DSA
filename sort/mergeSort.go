package sort

// import "fmt"

func merge(arr []int, low int, mid int, high int) {
	l := low
	r := mid + 1
	var temp []int
	for l <= mid && r <= high {
		if arr[l] <= arr[r] {
			temp = append(temp, arr[l])
			l++
		} else {
			temp = append(temp, arr[r])
			r++
		}
	}
	for l <= mid {
		temp = append(temp, arr[l])
		l++
	}
	for r <= high {
		temp = append(temp, arr[r])
		r++
	}
	for i := low; i <= high; i++ {
		arr[i] = temp[i - low]
	}
}

func MergeSort(arr []int, low int, high int) {

	if low >= high {
		return
	}
	mid := (low + high) / 2

	MergeSort(arr, low, mid)
	MergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
	
}