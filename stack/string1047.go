package main

//func removeDulplicates(S string) string{
//	stack := []byte{}
//	for i := range S{
//		if len(stack) > 0 && stack[len(stack) - 1] == S[i] {
//
//		}
//	}
//}


//使用栈的：

func foo1(S string) string {
		stack := []byte{}
		for i := range S{
			if len(stack) > 0 && stack[len(stack) - 1] == S[i] {
				//如果相等，出栈，
				stack = stack[:len(stack) - 1]
			} else {
				stack = append(stack, S[i])
			}
		}
		return string(stack)
}

//使用字符串转换数组的
var stack []byte = make([]byte,10000,10000)
var sArr []byte = make([]byte,10000,10000)

func helper(chars []byte, len int) {
	stack = stack[:0]


}

