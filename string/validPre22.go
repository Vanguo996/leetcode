package main

func isValid(s string) bool{
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}

	for i := 0; i < len(s); i++{
		// ？
		if m[s[i]] > 0 {
			//没有匹配的情况
			if len(stack) == 0 || stack[len(stack) - 1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack) - 1]  //栈弹出
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func isValid1(s string) bool{
	if len(s) == 0{
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s{
		if v == '[' || v =='{' || v == '(' {
			stack = append(stack, v)
		} else if (v == '}' && len(stack) > 0 && stack[len(stack) - 1] == '{') ||
			(v == ']' && len(stack) > 0 && stack[len(stack) - 1] == '[') ||
			(v == ')' && len(stack) > 0 && stack[len(stack) - 1] == '(') {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}

	}
	return len(stack) == 0
}


