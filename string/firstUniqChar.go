package main

func foo(s string) int {

	counts := make([]int, 26)

	for _, c := range s {
		//统计出现次数
		counts[c - 'a']++
	}

	for i, c := range s{
		if counts[c - 'a'] == 1 {
			return i
		}
	}
	return -1

}
