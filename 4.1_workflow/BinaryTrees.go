package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func build_tree() *Node {
	a_node := Node{"A", nil, nil}
	b_node := Node{"B", nil, nil}
	c_node := Node{"C", nil, nil}
	d_node := Node{"D", nil, nil}
	e_node := Node{"E", nil, nil}
	f_node := Node{"F", nil, nil}
	g_node := Node{"G", nil, nil}
	h_node := Node{"H", nil, nil}
	i_node := Node{"I", nil, nil}
	j_node := Node{"J", nil, nil}
	a_node.left = &b_node
	a_node.right = &c_node
	b_node.left = &d_node
	b_node.right = &e_node
	c_node.right = &f_node
	e_node.left = &g_node
	f_node.left = &h_node
	h_node.left = &i_node
	h_node.right = &j_node
	return &a_node
}

func (node *Node) display_indented(indent string, depth int) string {
	result := strings.Repeat(indent, depth) + node.data + "\n"
	if node.left != nil {
		result += node.left.display_indented(indent, depth+1)
	}
	if node.right != nil {
		result += node.right.display_indented(indent, depth+1)
	}
	return result
}

func (node *Node) preorder() string {
	result := node.data
	if node.left != nil {
		result += " " + node.left.preorder()
	}
	if node.right != nil {
		result += " " + node.right.preorder()
	}
	return result
}

func (node *Node) inorder() string {
	var result string

	if node.left != nil {
		result += node.left.inorder()
	}
	result += node.data + " "
	if node.right != nil {
		result += node.right.inorder()
	}

	return result
}

func (node *Node) postorder() string {
	var result string
	if node.left != nil {
		result += node.left.postorder()
	}
	if node.right != nil {
		result += node.right.postorder()
	}
	result += node.data + " "
	return result
}

func (node *Node) breadth_first() string {
	var result string
	return result
}

// *** DoublyLinkedList code ***
type Cell struct {
	data       *Node
	next, prev *Cell
}

type DoublyLinkedList struct {
	top_sentinel, bottom_sentinel *Cell
}

// Make a new DoublyLinkedList and initialize its sentinels.
func make_doubly_linked_list() DoublyLinkedList {
	list := DoublyLinkedList{}

	top_cell := Cell{nil, nil, nil}
	bottom_cell := Cell{nil, nil, nil}

	list.top_sentinel = &top_cell
	list.top_sentinel.next = &bottom_cell
	list.top_sentinel.prev = nil

	list.bottom_sentinel = &bottom_cell
	list.bottom_sentinel.prev = list.top_sentinel
	list.bottom_sentinel.next = nil
	return list
}

// Add a cell after me.
func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	after.prev = me

	me.next.prev = after
	me.next = after
}

// Add a cell before me.
func (me *Cell) add_before(before *Cell) {
	me.prev.add_after(before)
}

// Delete the cell and return it.
func (me *Cell) delete() *Cell {
	me.prev.next = me.next
	me.next.prev = me.prev
	return me
}

// Return true if the stack is empty, false otherwise.
func (stack *DoublyLinkedList) is_empty() bool {
	return stack.top_sentinel.next == stack.bottom_sentinel
}

// *** Stack functions ***

// Push an item onto the top of the list right after the sentinel.
func (stack *DoublyLinkedList) push(node *Node) {
	cell := Cell{data: node}
	stack.top_sentinel.add_after(&cell)
}

// Pop an item off of the list (from right after the sentinel).
func (stack *DoublyLinkedList) pop() *Node {
	if stack.is_empty() {
		panic("Cannot pop from an empty list")
	}

	return stack.top_sentinel.next.delete().data
}

// *** Queue functions ***

// Add an item to the top of the queue.
func (queue *DoublyLinkedList) enqueue(node *Node) {
	queue.push(node)
}

// Remove an item from the bottom of the queue.
func (queue *DoublyLinkedList) dequeue() *Node {
	if queue.is_empty() {
		panic("Cannot dequeue from an empty list")
	}
	return queue.bottom_sentinel.prev.delete().data
}

func main() {
	a_node := build_tree()
	fmt.Println(a_node.display_indented("  ", 0))
	fmt.Println("Preorder:     ", a_node.preorder())
	fmt.Println("Inorder:      ", a_node.inorder())
	fmt.Println("Postorder:    ", a_node.postorder())
	fmt.Println("Breadth first:", a_node.breadth_first())
}
