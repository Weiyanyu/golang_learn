package main

type tree struct {
	value       int
	left, right *tree
}

func TreeSort(values []int) {
	var root *tree
	for _, value := range values {
		root = add(root, value)
	}
	appendValues(root, values[:0])
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(t *tree, values []int) []int {
	if t != nil {
		values = appendValues(t.left, values)
		values = append(values, t.value)
		values = appendValues(t.right, values)
	}
	return values
}
