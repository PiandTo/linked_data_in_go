package main

import (
	"fmt"
	"bytes"
)

type Cell struct {
	date string
	next *Cell
	prev *Cell
}

type DoublyLinkedList struct {
	top_sentinel *Cell
	bottom_sentinel *Cell
}

func make_doubly_linked_list() *DoublyLinkedList {
	arr := DoublyLinkedList{ }
	arr.top_sentinel = &Cell {"TOP_CENTINEL", nil}
	arr.bottom_sentinel = &Cell {"BOTTOM_CENTINEL", nil}
	top_sentinel.next = bottom_sentinel
	bottom_sentinel.next = top_sentinel
	return &arr
}

func main() {
	
}