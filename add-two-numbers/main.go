package main

import (
	"math/big"
	"strconv"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := ListInt(l1), ListInt(l2)
	return ToList(num1.Add(num1, num2))
}

func ToList(num *big.Int) *ListNode {
	var (
		prevNode    *ListNode
		currentNode *ListNode
	)
	for _, char := range num.String() {
		val, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		currentNode = &ListNode{
			Val:  val,
			Next: prevNode,
		}
		prevNode = currentNode
	}
	return currentNode
}

func ListInt(n *ListNode) *big.Int {
	var (
		stack       = make([]int, 0)
		currentNode = n
	)
	for { // iterate through linked-list and push Val(s) onto stack
		stack = append(stack, currentNode.Val)
		if currentNode.Next == nil {
			break
		}
		currentNode = currentNode.Next
	}

	var str string
	for len(stack) > 0 { // pop Val(s) off stack and build string
		var i int
		i, stack = stack[len(stack)-1], stack[:len(stack)-1]
		str += strconv.Itoa(i)
	}

	bInt := new(big.Int)
	if _, ok := bInt.SetString(str, 10); !ok {
		panic("failed to bInt.SetString")
	}
	return bInt
}
