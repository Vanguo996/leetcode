package main

//判断链表是否有环
//快慢指针
func cycle(head *ListNode) bool{
	slow := head
	fast := head

	for fast != nil && slow != nil{
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}

	}
	return false

}
