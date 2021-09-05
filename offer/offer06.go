package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint1(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return recur(head)
}

func recur(head *ListNode) []int {
	if head.Next != nil {
		list := recur(head.Next)
		list = append(list, head.Val)
		return list
	}

	return []int{head.Val}
}

//新的递归写法：
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return []int{}  //返回空的切片
	}
	res := reversePrint2(head.Next)
	return append(res, head.Val)
}


//方法2：反转链表、再生成数组

func reversePrint3(head *ListNode) []int {

	if head == nil {
		return nil
	}
	//反转
	var pre *ListNode
	cur := head
	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	var res []int
	for pre != nil {
		res = append(res, pre.Val)
		pre = pre.Next
	}
	return res

}

//方法3 先顺序放入，再反转

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for i, j := 0, len(arr) - 1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

