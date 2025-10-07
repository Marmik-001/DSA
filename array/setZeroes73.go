package array

import "fmt"

func SetZeroes(m [][]int) {
	zeroes := [][]int{}
	for i, iv := range m {

		for j , jv := range iv {
			if jv == 0  {
				zeroes = append(zeroes, []int{i,j})
			}
		} 
	}

	for _ , v := range zeroes {
		m[v[0]] = make([]int, len(m[0]))
		for i := range m {
			m[i][v[1]] = 0
		}
	}
	fmt.Println(m)
}