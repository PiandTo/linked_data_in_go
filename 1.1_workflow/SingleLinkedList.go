package main

import (
	"fmt"
	"bytes"
)

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

func (list *LinkedList) print_linked_list() {
	for cell := list.sentinel; cell != nil; cell = cell.next {
		fmt.Printf("%s ", cell.data)
	}
}

func make_linked_list() LinkedList {
	arr := LinkedList{}
	arr.sentinel = &Cell {"CENTINEL", nil}
	return arr
}

func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	me.next = after
}

func (me *Cell) delete_after(after *Cell) (*Cell, error) {
	if after == nil {
		return nil, fmt.Errorf("No")
	}
	me.next = after.next
	after.next = me
	return after, nil
}

func (list *LinkedList) add_range(values []string) {
	var last *Cell
	for i := list.sentinel; i != nil; i = i.next {
		last = i
	}
	// fmt.Printf("%s ", last.data)
	for _, v := range values {
		// fmt.Printf("%d %s", i, v)
		a := Cell {v, nil}
		(last).add_after(&a)
		last = &a
	}
	// list.print_linked_list()
}

func (list *LinkedList) to_string(separator string) string {
	var strs bytes.Buffer
	for cell := list.sentinel; cell != nil; cell = cell.next {
		strs.WriteString(cell.data)
		// fmt.Printf("%s", cell.data)
		if cell.next != nil {
			strs.WriteString(separator)
			// fmt.Printf("%s ", separator)
		}
	}
	return strs.String()
}

func (list *LinkedList)length() int {
	i := -1
	for top := list.sentinel; top != nil; top = top.next {
		i++
	}
	return i
}

func (list *LinkedList)is_empty() bool {
	if list.sentinel.next == nil {
		return true
	} else {
		return false
	}

}

func (list *LinkedList) push(param string) {
	tmp := Cell {data: param}
	(list.sentinel).add_after(&tmp)
}

func (list *LinkedList)pop() string{
	cell, _ := (list.sentinel).delete_after(list.sentinel.next)
	return cell.data
}

func main() {
	// small_list_test()

   // Make a list from a slice of values.
   greek_letters := []string {
       "α", "β", "γ", "δ", "ε",
   }
   list := make_linked_list()
//    list1 := make_linked_list()
   list.add_range(greek_letters)
   fmt.Println(list.to_string(" "))
   fmt.Println()
//    length := list.length()
//    fmt.Printf("%d", length)
//    fmt.Println()
//    fmt.Printf("%t", list1.is_empty())
//    fmt.Println()
   // Demonstrate a stack.
   stack := make_linked_list()
   stack.push("Apple")
   stack.push("Banana")
   stack.push("Coconut")
   stack.push("Date")
   for !stack.is_empty() {
       fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
           stack.pop(),
           stack.length(),
           stack.to_string(" "))
   }
}