package twoPointers

//func lengthOfLongestSubstring(s string) int {

	//left, right := 0, -1
	//for left < len(s) {
	//	if right + 1 < len(s) && freq[s[right + 1] - 'a'] == 0 {
	//		freq[s[right + 1] - 'a']++
	//		right++
	//	} else {
	//		freq[s[left] - 'a']--
	//		left++
	//	}
	//}
	//res := max(res, left - right)










	////位图
	//if len(s) == 0 {
	//	return 0
	//}
	////扩展的ASCII码的位图表示,
	//var bitSet [256]uint8
	////m := make(map[byte]int)
	////m1 := map[byte]int{}
	//res, left, right := 0, 0, 0
	//for left < len(s) {
	//	if right < len(s) && bitSet[s[right]] == 0 {
	//		//标记对应点ASCII码为1
	//		bitSet[s[right]] = 1
	//		right++
	//	} else {
	//		//s[left] 表示字符串中的每一个字符，每个字符底层又是用byte存储，
	//		bitSet[s[left]] = 0
	//		left++
	//	}
	//	res = int(math.Max(float64(res), float64(right-left)))
	//}
	//return res

//}




