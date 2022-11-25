package leetcode

import "GoCodingProblems/model"

/*
*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

https://leetcode.com/problems/merge-two-sorted-lists/?envType=study-plan&id=level-1

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]

Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
Accepted
2.8M
Submissions
4.5M
Acceptance Rate
*/

func mergeTwoLists(list1 *model.ListNode, list2 *model.ListNode) *model.ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *model.ListNode
	if list1.Val <= list2.Val {
		head = list1
		head.Next = mergeTwoLists(list1.Next, list2)
	} else {
		head = list2
		head.Next = mergeTwoLists(list1, list2.Next)
	}

	return head
}
