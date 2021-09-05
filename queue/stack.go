package queue

import "container/list"

type CQueue struct {
	s1, s2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		s1: list.New(),
		s2: list.New(),
	}
}

func (this *CQueue) appendTail(value int)  {
	this.s1.PushBack(value)
}



