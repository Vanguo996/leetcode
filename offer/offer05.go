package main

import "fmt"

func replaceSpace(s string) string {

	fmt.Print([]byte(s))  // [37 50 48]
	fmt.Print([]rune(s))  // [37 50 48]
	res := ""
	//for _, v := range s {
	//
	//	if v == ' ' { //v 为int32类型
	//		res += "%20"
	//	} else {
	//		res += string(v)
	//	}
	//}
	return res
}


// 了解rune、byte区别：字符是rune类型，
// string底层是按字节存储的，一个字符可能是任意字节，rune一个字符是4字节

func replaceSpace2(s string) string {
	var b []rune   // rune是int32的别名
	for _, v := range s {
		if v == 32 {
			b = append(b, 37, 50, 48)
		} else {
			b = append(b, v)
		}
	}
	return string(b)
}
