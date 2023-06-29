package main

import (
	"bytes"
	"fmt"
)

type Cell struct {
	date string
	next *Cell
	prev *Cell
}

type DoublyLinkedList struct {
	top_sentinel    *Cell
	bottom_sentinel *Cell
}

func make_doubly_linked_list() *DoublyLinkedList {
	arr := DoublyLinkedList{}
	arr.top_sentinel = &Cell{"TOP_CENTINEL", nil, nil}
	arr.bottom_sentinel = &Cell{"BOTTOM_CENTINEL", nil, nil}
	arr.top_sentinel.next = arr.bottom_sentinel
	arr.bottom_sentinel.prev = arr.top_sentinel
	return &arr
}

func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	after.prev = me.prev
	// me.prev = after
}

func (me *Cell) add_before(before *Cell) {
	before.next = me
	before.prev = me.prev
	// fmt.Println(before.prev.date)
	// fmt.Println(before.next.date)
	me.prev = before

	// fmt.Println(before.prev.date)
	// fmt.Println(before.date)
	// fmt.Println(before.next.date)
	// fmt.Println("-----------")
}

func (list *DoublyLinkedList) print_linked_list() {
	for cell := list.top_sentinel; cell != list.bottom_sentinel; cell = cell.next {
		fmt.Printf("%s ", cell.date)
	}
}

func delete(me *Cell) {

}

func (list *DoublyLinkedList) add_range(values []string) {
	// var last *Cell
	// for i := list.top_sentinel; i != nil; i = i.next {
	// 	last = i
	// }
	// fmt.Printf("%s ", last.data)
	for _, v := range values {
		// fmt.Printf("%s", v)
		a := &Cell{v, nil, nil}
		(list.bottom_sentinel).add_before(a)
		// fmt.Println(list.top_sentinel.date)
		// list.bottom_sentinel = &a
	}
	// list.print_linked_list()
}

func (list *DoublyLinkedList) to_string(separator string) string {
	var strs bytes.Buffer

	// fmt.Printf("%s", list.top_sentinel.next.date)
	for cell := list.top_sentinel; cell != nil; cell = cell.next {
		// fmt.Println(cell.date)
		strs.WriteString(cell.date)
		// fmt.Printf("%s", cell.date)
		if cell.next != nil {
			strs.WriteString(separator)
			// 		// fmt.Printf("%s ", separator)
		}
	}
	return strs.String()
}

func main() {
	// Make a list from a slice of values.
	list := make_doubly_linked_list()
	animals := []string{
		"Ant",
		"Bat",
		"Cat",
		"Dog",
		"Elk",
		"Fox",
	}
	list.add_range(animals)
	fmt.Println(list.to_string(" "))
}
