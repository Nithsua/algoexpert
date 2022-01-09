package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	registry := map[int]bool{}
	temp := linkedList
	var prev *LinkedList = nil
	for temp != nil {
		if _, ok := registry[temp.Value]; ok {
			prev.Next = temp.Next
		} else {
			registry[temp.Value] = true
			prev = temp
		}
		temp = temp.Next
	}
	return linkedList
}
