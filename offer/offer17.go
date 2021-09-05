package main

func printNumbers(n int) []int {
	j := 1
	for i := 1; i <= n; i++ {
		j *= 10
	}

	var res []int
	for k := 1; k <= j - 1; k++ {
		res = append(res, k)
	}
	return res

}
