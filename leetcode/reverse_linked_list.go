package leetcode

import (
	"GoCodingProblems/model"
)

/*
*

Given the head of a singly linked list, reverse the list, and return the reversed list.

Constraints:

The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000

Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?

https://leetcode.com/problems/reverse-linked-list/?envType=study-plan&id=level-1
*/
func reverseList(head *model.ListNode) *model.ListNode {

	current := head
	var newHead *model.ListNode
	for current != nil {
		temp := current
		current = current.Next
		temp.Next = newHead
		newHead = temp
	}
	return newHead
}
