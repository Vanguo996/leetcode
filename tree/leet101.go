package tree


func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil{
		return true
	}
	if p == nil || q == nil {
		return false
	}
	// p指针右移动，q指针左移动
	check(p.Right, q.Left)
	//p指针左移动，q指针右移动，
	check(p.Left, q.Right)
	//每一个点需要相等，否则返回false
	return p.Val == q.Val


	//需要以下语句才正确：
	//return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)

}


//迭 代方式， 使用队列实现
func check1(root *TreeNode) bool {
	t1, t2 := root, root
	q := []*TreeNode{}
	q = append(q, t1)
	q = append(q, t2)

	for len(q) > 0 {
		t1, t2 = q[0], q[1]
		q = q[2:]
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		q = append(q, t1.Left)
		q = append(q, t2.Right)

		q = append(q, t1.Right)
		q = append(q, t2.Left)
	}
	return true
}