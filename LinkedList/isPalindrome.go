package main
type ListNode1 struct {
	Val int
	Next *ListNode1

}
func isPalindrome(head *ListNode1) bool{
	//先把链表值赋值到数组中
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l, r :=  0, len(arr) - 1
	for l < r{
		if arr[l] == arr[r]{
			l++
			r--
		} else{
			return false
		}
	}
	return true
}

func isPalindrome1(head *ListNode1) bool {
	//快慢指针
	slow, fast := head, head
	//边界条件
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//慢指针到中间，开始反转链表
	cur := slow.Next
	for cur != nil {
		tmp := slow.Next.Next
		cur.Next = slow //反转
		slow = cur
		cur = tmp
	}

	//反转完成，开始判断回文,
	return false

}


// 官方题解

//func reverseList(head *ListNode) *ListNode {
//	var prev, cur *ListNode = nil, head
//	for cur != nil {
//		nextTmp := cur.Next
//		cur.Next = prev
//		prev = cur
//		cur = nextTmp
//	}
//	return prev
//}
//
//func endOfFirstHalf(head *ListNode) *ListNode {
//	fast := head
//	slow := head
//	for fast.Next != nil && fast.Next.Next != nil {
//		fast = fast.Next.Next
//		slow = slow.Next
//	}
//	return slow
//}
//
//func isPalindrome(head *ListNode) bool {
//	if head == nil {
//		return true
//	}
//
//	// 找到前半部分链表的尾节点并反转后半部分链表
//	firstHalfEnd := endOfFirstHalf(head)
//	secondHalfStart := reverseList(firstHalfEnd.Next)
//
//	// 判断是否回文
//	p1 := head
//	p2 := secondHalfStart
//	result := true
//	for result && p2 != nil {
//		if p1.Val != p2.Val {
//			result = false
//		}
//		p1 = p1.Next
//		p2 = p2.Next
//	}
//
//	// 还原链表并返回结果
//	firstHalfEnd.Next = reverseList(secondHalfStart)
//	return result
//}















