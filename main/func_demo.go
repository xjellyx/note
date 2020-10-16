package main

import (
	"container/list"
	"fmt"
)

func main() {
	//fmt.Println(maxDepth(&Node{
	//	Val: 1,
	//	Children: []*Node{
	//		{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}},
	//		{Val: 2},
	//		{Val: 4},
	//	},
	//}))
	fmt.Println()
	fmt.Println(
		flatten(&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
				//Right: &TreeNode{
				//	Val:   5,
				//	Left:  nil,
				//	Right: nil,
				//},

			},
			Right: &TreeNode{Val: 1},
			//Right: &TreeNode{
			//	Val: 20,
			//	Left: &TreeNode{
			//		Val:   15,
			//		Left:  nil,
			//		Right: nil,
			//	},
			//	Right: &TreeNode{
			//		Val:   7,
			//		Left:  nil,
			//		Right: nil,
			//	},
			//},
		}))

}

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *Node) [][]int {
	var (
		queue = list.New()
		data  [][]int
	)
	if root == nil {
		return [][]int{}
	}
	queue.PushFront(root)
	for queue.Len() > 0 {
		count := queue.Len()
		arr := []int{}
		for count > 0 {
			element := queue.Back()
			node := element.Value.(*Node)
			arr = append(arr, node.Val)
			queue.Remove(element)
			if len(node.Children) > 0 {
				for _, v := range node.Children {
					queue.PushFront(v)
				}
			}
			count--
		}
		data = append(data, arr)
	}
	return data
}
func bsf(root *TreeNode, m *[]int) {
	if root == nil {
		return
	}
	*m = append(*m, root.Val)
	bsf(root.Left, m)
	bsf(root.Right, m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
