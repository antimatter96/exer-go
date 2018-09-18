package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Build returns a flattened array from an array of nodeId, parentId
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	var hasRoot bool
	var maxIdAllowed = len(records)
	var maxId int = 0

	grid := make([][]int, maxIdAllowed)

	for _, record := range records {
		if record.ID == 0 {
			if hasRoot {
				return nil, fmt.Errorf("Duplicate root")
			}
			hasRoot = true
			if record.Parent != 0 {
				return nil, fmt.Errorf("Root has parent")
			}
			continue
		}
		if record.ID >= maxIdAllowed || record.Parent >= maxIdAllowed {
			return nil, fmt.Errorf("Non continous")
		}
		if record.ID <= record.Parent {
			return nil, fmt.Errorf("Parent is smaller than child")
		}
		grid[record.Parent] = append(grid[record.Parent], record.ID)
		if record.ID > maxId {
			maxId = record.ID
		}
	}
	if !hasRoot {
		return nil, fmt.Errorf("Does not have root")
	}
	if hasCycle(grid) {
		return nil, fmt.Errorf("Cycle")
	}

	for _, v := range grid {
		sort.Ints(v)
	}

	mp := make(map[int]*Node)
	for i := maxId; i > -1; i-- {
		var node Node
		node.ID = i
		if len(grid[i]) > 0 {
			node.Children = make([]*Node, 0)
			for _, c := range grid[i] {
				node.Children = append(node.Children, mp[c])
			}
		}
		mp[i] = &node
	}
	return mp[0], nil
}

func hasCycle(grid [][]int) bool {
	checked := make(map[int]bool)
	stack := newStack()
	stack.push(0)
	for stack.len() != 0 {
		id := stack.pop()
		if id != -1 {
			if checked[id] {
				return true
			}
			checked[id] = true
			if len(grid[id]) == 0 {
				continue
			}
			for _, child := range grid[id] {
				stack.push(child)
			}
		}
	}
	return false
}

type (
	stack struct {
		top    *stackNode
		length int
	}
	stackNode struct {
		value int
		prev  *stackNode
	}
)

// Create a new stack
func newStack() *stack {
	return &stack{nil, 0}
}

// Return the number of items in the stack
func (this *stack) len() int {
	return this.length
}

// View the top item on the stack
func (this *stack) peek() int {
	if this.length == 0 {
		return -1
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *stack) pop() int {
	if this.length == 0 {
		return -1
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *stack) push(value int) {
	n := &stackNode{value, this.top}
	this.top = n
	this.length++
}
