package main

func lengthOfLongestSubstring(s string) int {
	res, left, right := 0, 0, -1
	var m [256]int

	for left < len(s) {
		if right + 1 < len(s) && m[s[right + 1] - 'a'] == 0{
			m[s[right+1] - 'a']++
			right++
		} else {
			m[s[left] - 'a']--
			left++
		}
		res = max(res, right - left + 1)
	}
	return res
}

func foo(s string) {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	//
	right, left := -1, 0

	for i := 0; i < len(s); i++{

		if i != 0 {
			// 左指针向右，删除字符
			delete(m, s[i-1])
		}

		// 字符没有在集合中，向右移动
		for right + 1 < len(s) && m[s[right + 1]] == 0 {
			m[s[right+1]]++
			right++
		}

		left = max(left, right - i - 1)

	}


}

func max(x, y int) int{
	if x < y {
		return y
	} else {
		return x
	}
}