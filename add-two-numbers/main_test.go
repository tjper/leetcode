package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		Name     string
		L1       *ListNode
		L2       *ListNode
		Expected int64
	}{
		{
			Name:     "Simple",
			L1:       NewNodeList(2, 4, 3),
			L2:       NewNodeList(5, 6, 4),
			Expected: 807,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			L3 := addTwoNumbers(test.L1, test.L2)
			actual := ListInt(L3)
			assert.Equal(t, test.Expected, actual.Int64())
		})
	}
}

func NewNodeList(values ...int) *ListNode {
	var (
		nextNode    *ListNode
		currentNode *ListNode
	)

	for i := len(values) - 1; i >= 0; i-- {
		currentNode = &ListNode{
			Val:  values[i],
			Next: nextNode,
		}
		nextNode = currentNode
	}
	return currentNode
}
