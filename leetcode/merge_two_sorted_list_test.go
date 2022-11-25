package leetcode

import (
	"GoCodingProblems/list_utils"
	"GoCodingProblems/model"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {

	t.Run("one element list with same value", func(t *testing.T) {
		list1 := &model.ListNode{Val: 1}
		list2 := &model.ListNode{Val: 1}

		got := mergeTwoLists(list1, list2)
		wantHead := list1
		wantArr := []int{1, 1}

		list_utils.AssertMergeList(t, got, wantArr, wantHead)

	})

	t.Run("one element list where list2 has smaller value", func(t *testing.T) {
		list1 := &model.ListNode{Val: 2}
		list2 := &model.ListNode{Val: 1}

		got := mergeTwoLists(list1, list2)
		wantHead := list2
		wantArr := []int{1, 2}

		list_utils.AssertMergeList(t, got, wantArr, wantHead)
	})

	t.Run("mixed values in two lists", func(t *testing.T) {
		list1 := &model.ListNode{Val: 2, Next: &model.ListNode{Val: 7}}
		list2 := &model.ListNode{Val: 3}

		got := mergeTwoLists(list1, list2)
		wantHead := list1
		wantArr := []int{2, 3, 7}

		list_utils.AssertMergeList(t, got, wantArr, wantHead)
	})

}
