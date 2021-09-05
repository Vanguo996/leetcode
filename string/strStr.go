package main

func strStr(haystack string, needle string) int {
	// 在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
	// 对于本题而言，当 needle 是空字符串时我们应当返回 0

	for i := 0; ; i++{
		for j := 0; ; j++{
			if j == len(needle){
				return i
			}
			if j + i == len(haystack){
				return -1
			}
			if needle[j] != haystack[i+j]{
				break
			}

		}
	}


}
