package leetcode

import (
	"GoCodingProblems/list_utils"
	"GoCodingProblems/model"
	"testing"
)

func Test_reverseList(t *testing.T) {
	t.Run("one element list with same value", func(t *testing.T) {
		list1 := &model.ListNode{Val: 1}

		got := reverseList(list1)
		wantHead := list1
		wantArr := []int{1}

		list_utils.AssertMergeList(t, got, wantArr, wantHead)
	})

	t.Run("two elements list with same value", func(t *testing.T) {
		tail := &model.ListNode{Val: 2}
		head := &model.ListNode{Val: 1, Next: tail}

		got := reverseList(head)
		wantHead := tail
		wantArr := []int{2, 1}

		list_utils.AssertMergeList(t, got, wantArr, wantHead)
	})

	t.Run("three element list with same value", func(t *testing.T) {
		tail := &model.ListNode{Val: 3}
		mid := &model.ListNode{Val: 2, Next: tail}
		head := &model.ListNode{Val: 1, Next: mid}

		got := reverseList(head)
		wantHead := tail
		wantArr := []int{3, 2, 1}

		list_utils.AssertMergeList(t, got, wantArr, wantHead)
	})

	//t.Run("mixed values in two lists", func(t *testing.T) {
	//	list1 := &model.ListNode{Val: 2, Next: &model.ListNode{Val: 7}}
	//	list2 := &model.ListNode{Val: 3}
	//
	//	got := mergeTwoLists(list1, list2)
	//	wantHead := list1
	//	wantArr := []int{2, 3, 7}
	//
	//	list_utils.AssertMergeList(t, got, wantArr, wantHead)
	//})

}
