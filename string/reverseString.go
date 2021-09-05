package main



func reverse(s []byte) []byte {
	for i := 0; i < len(s) / 2; i++{
		s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
	}
	return s
}

// 345  反转元音字母，对撞指针

func reverseVowel(s string) string {
	b := []byte(s)
	for l , r := 0, len(s); l < r ;{
		for !verify(b[l]) {
			l++
		}
		for l < r && !verify(b[r]) {
			r--
		}
		b[r], b[l] = b[l], b[r]
		l++
		r--
	}
	return string(b)
}

func verify(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' ||
		s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}


