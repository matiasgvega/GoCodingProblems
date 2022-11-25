package list_utils

import (
	"GoCodingProblems/model"
	"reflect"
	"testing"
)

func GetValues(head *model.ListNode) (list []int) {
	list = []int{head.Val}
	for head.Next != nil {
		head = head.Next
		list = append(list, head.Val)
	}
	return
}

func AssertMergeList(t *testing.T, got *model.ListNode, wantArr []int, wantHead *model.ListNode) {
	t.Helper()

	if got != wantHead {
		t.Errorf("mergeTwoLists() = %v, wantHead %v", got, wantHead)
	}

	gotArr := GetValues(got)
	if !reflect.DeepEqual(gotArr, wantArr) {
		t.Errorf("getValues = %v, wantArr %v", gotArr, wantArr)
	}
}
