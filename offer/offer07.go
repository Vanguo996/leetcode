package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	//
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	//找到分割点
	i := 0
	for ; i < len(inorder) - 1; i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	//左子树
	root.Left = buildTree(preorder[1:len(inorder[:i]) + 1],inorder[:i])
	//右子树
	root.Right = buildTree(preorder[len(inorder[:i]) + 1:], inorder[i + 1:])

	return root

}
