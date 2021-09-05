package main



func reversePrint(head *ListNode) []int {

	if head == nil {return nil}

	return recur(head)

}

func recur(head *ListNode) []int {

	if head.next != nil {
		list := recur(head.next)
		list = append(list, head.val)
		return list
	}

	return []int{head.val}
}




