package main

import (
	"strings"
	"unicode"
)

func main() {

	numDifferentIntegers("aaa23dfg2345")

}


func numDifferentIntegers(word string) int {
	//超时：
	//set := make(map[string]bool)
	//i := 0
	//for i < len(word) {
	//	if unicode.IsNumber(rune(word[i])) {
	//		//检测有多少数字
	//		r := i
	//		for r + 1 < len(word) && unicode.IsNumber(rune(word[r + 1])) {
	//			set[word[i : r + 1]] = true
	//		}
	//		i = r + 1
	//	} else {
	//		i++
	//	}
	//
	//}
	//return len(set)

	set := map[string]struct{}{}
	//分割小写字母，
	for _, s := range strings.FieldsFunc(word, unicode.IsLower) {
		for len(s) > 1 && s[0] == '0' {
			s = s[1:]
		}
		set[s] = struct{}{}
	}
	return len(set)


}
