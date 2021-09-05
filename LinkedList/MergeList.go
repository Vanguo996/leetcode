package main

//合并两个有序链表


//func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
//	设定哨兵节点
//	preHead := &ListNode{}
//	result := preHead
//
//	if l1.val < l2.val{
//		preHead
//	} else {
//
//	}
//
//}

func re(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}

	if l1.val < l2.val {
		l1.next = re(l1.next, l2)
		return l1
	}
	l2.next = re(l1, l2.next)
	return l2

}



