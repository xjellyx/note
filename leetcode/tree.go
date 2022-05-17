package main

import (
	"container/list"
	"fmt"
	"github.com/olongfen/note/leetcode/utils"
	"math"
)

func t1(tree *utils.TreeNode) (res []int) {
	var (
		stack = list.New()
	)
	if tree == nil {
		return
	}

	stack.PushFront(tree)

	for stack.Len() > 0 {
		var (
			e    = stack.Front()
			node = e.Value.(*utils.TreeNode)
		)
		stack.Remove(e)
		res = append(res, node.Val)
		if node.Right != nil {
			stack.PushFront(node.Right)
		}
		if node.Left != nil {
			stack.PushFront(node.Left)
		}
	}
	return
}

func t2(root *utils.TreeNode) (res []int) {
	var (
		stack = list.New()
	)
	cur := root
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushFront(cur)
			cur = cur.Left
		} else {
			e := stack.Front()
			stack.Remove(e)
			cur = e.Value.(*utils.TreeNode)
			res = append(res, cur.Val)
			cur = cur.Right
		}

	}

	return
}

func t3(tree *utils.TreeNode) (res []int) {
	var (
		stack = list.New()
	)
	if tree == nil {
		return
	}

	stack.PushFront(tree)

	for stack.Len() > 0 {
		var (
			e    = stack.Front()
			node = e.Value.(*utils.TreeNode)
		)
		stack.Remove(e)
		res = append(res, node.Val)
		if node.Left != nil {
			stack.PushFront(node.Left)
		}
		if node.Right != nil {
			stack.PushFront(node.Right)
		}
	}
	reverseArr(res)
	return
}

func reverseArr(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func all(root *utils.TreeNode) (res []int) {
	if root == nil {
		return
	}
	var (
		stack = list.New()
		node  *utils.TreeNode
	)
	stack.PushFront(root)
	for stack.Len() > 0 {
		e := stack.Front()
		stack.Remove(e)
		if e.Value == nil {
			e = stack.Front()
			stack.Remove(e)
			node = e.Value.(*utils.TreeNode)
			res = append(res, node.Val)
			continue
		}

		node = e.Value.(*utils.TreeNode)
		// 前序遍历：中左右
		// 压栈顺序：右左中
		if node.Right != nil {
			stack.PushFront(node.Right)
		}
		if node.Left != nil {
			stack.PushFront(node.Left)
		}
		stack.PushFront(node)
		stack.PushFront(nil)
		// 中序遍历: 左中右
		// 压栈顺序: 右中左
		/*if node.Right != nil {
			stack.PushFront(node.Right)
		}
		stack.PushFront(node)
		stack.PushFront(nil)
		if node.Left != nil {
			stack.PushFront(node.Left)
		}*/
		// 后序遍历: 左右中
		// 压栈顺序： 中右左
		/*	stack.PushFront(node)
			stack.PushFront(nil)
			if node.Right != nil {
				stack.PushFront(node.Right)
			}
			if node.Left != nil {
				stack.PushFront(node.Left)
			}*/
	}

	return
}

func c(root *utils.TreeNode) (res [][]int) {
	var (
		queue = list.New()
	)
	if root == nil {
		return
	}
	queue.PushFront(root)
	for queue.Len() > 0 {
		size := queue.Len()
		var tmpArr []int
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Back()).(*utils.TreeNode)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
			tmpArr = append(tmpArr, node.Val)
		}
		res = append(res, tmpArr)
	}
	return
}

func sumOfLeftLeaves(root *utils.TreeNode) int {
	var (
		stack = list.New()
		res   = 0
	)
	if root == nil {
		return 0
	}
	var (
		node *utils.TreeNode
	)
	stack.PushFront(root)
	for stack.Len() > 0 {
		e := stack.Front()
		stack.Remove(e)
		if e.Value == nil {
			e = stack.Front()
			stack.Remove(e)
			node = e.Value.(*utils.TreeNode)
			if node.Left == nil && node.Right == nil {
				res += node.Val
			}
			continue
		}

		node = e.Value.(*utils.TreeNode)
		// 前序遍历：中左右
		// 压栈顺序：右左中
		if node.Right != nil {
			stack.PushFront(node.Right)
		}
		if node.Left != nil {
			stack.PushFront(node.Left)
		}
		stack.PushFront(node)
		stack.PushFront(nil)
	}
	return res
}

func pathSum(root *utils.TreeNode, targetSum int) [][]int {

	if root == nil {
		return res
	}
	tmpArr = append(tmpArr, root.Val)
	traversal(root, targetSum-root.Val)
	return res
}

var (
	res    [][]int
	tmpArr []int
)

func traversal(node *utils.TreeNode, count int) {
	if node.Left == nil && node.Right == nil && count == 0 {
		var (
			data = make([]int, len(tmpArr))
		)
		copy(data, tmpArr)
		res = append(res, data)
		return
	}
	if node.Left == nil && node.Right == nil {
		return
	}

	if node.Left != nil {
		tmpArr = append(tmpArr, node.Left.Val)
		count -= node.Left.Val
		traversal(node.Left, count)
		count += node.Left.Val
		tmpArr = tmpArr[:len(tmpArr)-1]

	}

	if node.Right != nil {
		tmpArr = append(tmpArr, node.Right.Val)
		count -= node.Right.Val
		fmt.Println(tmpArr, count)
		traversal(node.Right, count)
		count += node.Right.Val
		tmpArr = tmpArr[:len(tmpArr)-1]
		fmt.Println("A", tmpArr, count)
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	left := findRootIndex(inorder, rootVal)
	fmt.Println(rootVal, left, preorder, inorder)
	root := &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:left+1], inorder[:left]),
		Right: buildTree(preorder[left+1:], inorder[left+1:]),
	}
	return root
}

func findRootIndex(inorder []int, target int) (index int) {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {

	node := &TreeNode{}
	build(root, node)
	return node.Right
}

func build(root *TreeNode, node *TreeNode) {
	if root == nil {
		return
	}
	build(root.Left, node)
	node.Right = &TreeNode{Val: root.Val}
	node = node.Right
	build(root.Right, node)

}

func getMinimumDifference(root *TreeNode) int {
	var (
		m   = 1 << 32
		pre *TreeNode
		fc  func(node *TreeNode)
	)
	fc = func(node *TreeNode) {
		if node == nil {
			return
		}
		fc(node.Left)
		if pre != nil {
			println("wwwwwwwwwwwwwwwwwww")
			d := int(math.Abs(float64(node.Val - pre.Val)))
			if m > d {
				m = d
			}
		}
		pre = node
		fc(node.Right)
	}
	fc(root)
	return m
}
func main() {
	t := &TreeNode{Val: 1, Left: &TreeNode{Val: 2,
		Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5},
	}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}

	_ = t
	fmt.Println(getMinimumDifference(&TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}}))
}
