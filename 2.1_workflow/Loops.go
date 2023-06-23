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

func (list LinkedList)has_loop() bool {
	slow := list.sentinel
	fast := list.sentinel.next
	for fast != slow {
		if fast == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
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

func (list LinkedList)to_string_max(separator string, max int) string {
	i := 0
	var strs bytes.Buffer
	for cell := list.sentinel; cell != nil; cell = cell.next {
		i++
		if i == max {
			break 
		}
		strs.WriteString(cell.data)
		// fmt.Printf("%s", cell.data)
		if cell.next != nil {
			strs.WriteString(separator)
			// fmt.Printf("%s ", separator)
		}
	}
	return strs.String()
}

func make_linked_list() LinkedList {
	arr := LinkedList{}
	arr.sentinel = &Cell {"CENTINEL", nil}
	return arr
}

func main() {
    // Make a list from a slice of values.
    values := []string {
        "0", "1", "2", "3", "4", "5",
    }
    list := make_linked_list()
    list.add_range(values)

    fmt.Println(list.to_string(" "))
    if list.has_loop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 5 point to cell 2.
    list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

    fmt.Println(list.to_string_max(" ", 10))
    if list.has_loop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 4 point to cell 2.
    list.sentinel.next.next.next.next.next = list.sentinel.next.next

    fmt.Println(list.to_string_max(" ", 10))
    if list.has_loop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
	}
}