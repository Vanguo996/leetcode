package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func travel(root *TreeNode) int{
	var res []int
	var nodes []*TreeNode = []*TreeNode{root}
	sum := 0
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		res = append(res, node.Val)
		sum += node.Val
		if (node.Left != nil) {
			nodes = append(nodes, node.Left)
		}
		if (node.Right != nil) {
			nodes = append(nodes, node.Right)
		}
	}
	return sum

}


func main()  {

}


