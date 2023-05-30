package main

import "fmt"

func nextLargerNodes(head *ListNode) (ans []int) {
	st := []int{} // 单调栈（节点下标）
	for cur := head; cur != nil; cur = cur.Next {
		x := cur.Val
		for len(st) > 0 && ans[st[len(st)-1]] < x {
			ans[st[len(st)-1]] = x // ans[st[len(st)-1]] 后面不会再访问到了
			st = st[:len(st)-1]
		}
		st = append(st, len(ans)) // 当前 ans 的长度就是当前节点的下标
		ans = append(ans, x)
	}
	for _, i := range st {
		ans[i] = 0 // 没有下一个更大元素
	}
	return ans
}

func main() {
	fmt.Println(nextLargerNodes(&ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5}}}))
}
