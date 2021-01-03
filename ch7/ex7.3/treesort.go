package main

import (
	"fmt"
	"bytes"
)

type tree struct {
	value       int
	left, right *tree
}

//Sort sorts values in place
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

//appendValues appends the element of t to values in order
func appendValues(values []int, t *tree) []int{
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}

	if t.value > v {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}

	return t
}

func (t *tree) String() string {
	values := make([]int, 0)
	values = appendValues(values, t)
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, value := range values {
		if i > 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", value)
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	root := &tree{value: 3}
	root = add(root, 2)
	root = add(root, 4)
	fmt.Println(root)
}