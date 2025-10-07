package array

func CountPrime(n int) int {

	arr := make([]bool, n)

	if n <= 1 {
		return 0
	}
	ans := 1
    if n == 2{
        return 0
    }
	for i:= range arr {
		if i <= 1 {
			arr[i] = false
		}else{
			arr[i] = true
		}
	}

	for i := 3; i < n; i=i+2 {
		if arr[i]{
			ans++
		}
		j := i*i
		for j < n {
			arr[j] = false
			j = j + i
		}

	}
	return ans
}