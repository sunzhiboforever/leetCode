package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	addTwoNumbers(l1, l2)
}

//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var last *ListNode
	up := 0
	for {
		if l1 == nil && l2 == nil && up == 0 {
			break
		}
		cur := &ListNode{
			Val:  0,
			Next: nil,
		}
		sum := up
		if l1 == nil && l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if l2 == nil && l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l1 != nil && l2 != nil {
			sum = sum + l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}
		if sum >= 10 {
			up = 1
			cur.Val = sum / 10
		} else {
			up = 0
			cur.Val = sum
		}
		if last != nil {
			last.Next = cur
		}
		last = cur
		if head == nil {
			head = cur
		}
	}
	return head
}
