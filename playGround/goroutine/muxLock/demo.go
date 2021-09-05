package muxLock

//func main() {
//	var wg sync.waitGroup
//
//
//}

type words struct {
	found map[string]int

}

//创建构造函数
func newWords() *words {
	return &words{found: map[string]int{}}
}

//找出单词出现次数
func (w words) add(word string, n int) { //跟踪所发现的单词次数
	count, ok := w.found[word]
	if !ok{
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}




