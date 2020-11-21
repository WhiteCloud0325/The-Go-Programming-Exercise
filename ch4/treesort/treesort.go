package main

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

func appendValues(values []int, t *tree) {
	if t == nil {
		return
	}

	appendValues(values, t.left)

	values = append(values, t.value)

	appendValues(values, t.right)
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
}
