package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var (
		dummyHead = &ListNode{}
	)

	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}

type MyLinkedList struct {
	ListNode *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {

	return MyLinkedList{ListNode: &ListNode{Val: -1}}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	head := this.ListNode.Next

	for head != nil && index > 0 {
		index--
		head = head.Next
	}
	if index != 0 {
		return -1
	}

	if head == nil {
		return -1
	}
	return head.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	var (
		head = &ListNode{Val: val}
	)
	head.Next = this.ListNode.Next
	this.ListNode.Next = head
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	var (
		tail = &ListNode{Val: val}
	)
	cur := this.ListNode.Next
	if cur == nil {
		this.ListNode.Next = tail
		return
	}
	for cur != nil {
		if cur.Next == nil {
			cur.Next = tail
			break
		}
		cur = cur.Next
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

	cur := this.ListNode.Next
	insert := cur
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	for index > 0 && cur != nil {
		insert = cur
		cur = cur.Next
		index--
	}
	if index < 0 {
		this.AddAtHead(val)
	} else if index == 0 && cur == nil {
		this.AddAtTail(val)
	} else if index == 0 && insert != nil {
		temp := insert.Next
		insert.Next = &ListNode{Val: val, Next: temp}
	}

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	del := this.ListNode
	cur := this.ListNode.Next
	for index > 0 && cur != nil {
		del = cur
		cur = cur.Next
		index--
	}
	if index == 0 && cur != nil {
		del.Next = cur.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil && head.Next != nil {
		fmt.Println(pre.Next, head)
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		pre = head
		head = next
		fmt.Println("www", pre.Next, head)
	}
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = reverseList(head)
	dummu := &ListNode{Next: head, Val: -1}
	prev := dummu.Next
	cur := dummu.Next
	for n > 0 && cur != nil {
		prev = cur
		cur = cur.Next
		n--
	}
	if n == 0 && cur != nil {
		prev.Next = cur.Next
	}
	return dummu.Next
}

func detectCycle(head *ListNode) *ListNode {
	var (
		dummy          = &ListNode{Val: -1, Next: head}
		fast           = head
		slow           = head
		index1, index2 *ListNode
	)
	for slow != nil && slow.Next != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			index1 = fast
			index2 = dummy.Next
			break
		}
	}

	for index1 != nil && index2 != nil && index1 != index2 {
		index1 = index1.Next
		index2 = index2.Next

		if index1 == index2 {
			return index1
		}
	}
	if index1 == index2 {
		return index1
	}
	return nil
}

func hasCycle(head *ListNode) bool {
	var (
		dummy = &ListNode{Next: head}
		left  = dummy.Next
		right = dummy.Next
	)
	cur := dummy.Next
	for cur != nil && cur.Next != nil {
		left = cur
		right = cur.Next.Next
		if left == right {
			return true
		}
		cur = cur.Next
	}
	return false
}

func main() {
	_t := &ListNode{Val: 2}
	head := &ListNode{Val: 3, Next: _t}

	_t2 := &ListNode{Val: 4, Next: nil}
	_t1 := &ListNode{Val: 0, Next: _t2}
	_t.Next = _t1

	//_t.Next = _t1
	fmt.Println("aa", hasCycle(head))
}
