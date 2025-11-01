package main

// go test -v homework_test.go

type node struct {
	key, value  int
	left, right *node
}

type OrderedMap struct {
	root *node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{} // need to implement
}

func (m *OrderedMap) Insert(key, value int) {
	m.root = insertNode(m.root, key, value, &m.size)
}

func insertNode(n *node, key, value int, size *int) *node {
	if n == nil {
		*size++
		return &node{key: key, value: value}
	}

	if key < n.key {
		n.left = insertNode(n.left, key, value, size)
	} else if key > n.key {
		n.right = insertNode(n.right, key, value, size)
	} else {
		n.value = value // update value if key already exists
	}

	return n
}

func (m *OrderedMap) Erase(key int) {
	var erased bool
	m.root, erased = eraseNode(m.root, key)
	if erased {
		m.size--
	}
}

func eraseNode(n *node, key int) (*node, bool) {
	if n == nil {
		return nil, false
	}

	if key < n.key {
		var deleted bool
		n.left, deleted = eraseNode(n.left, key)
		return n, deleted
	}

	if key > n.key {
		var deleted bool
		n.right, deleted = eraseNode(n.right, key)
		return n, deleted
	}

	// key == n.key → удаляем этот узел
	if n.left == nil {
		return n.right, true
	}
	if n.right == nil {
		return n.left, true
	}

	// два потомка
	successor := findMin(n.right)
	n.key, n.value = successor.key, successor.value
	var deleted bool
	n.right, deleted = eraseNode(n.right, successor.key)
	return n, deleted
}

func findMin(n *node) *node {
	for n.left != nil {
		n = n.left
	}

	return n
}

func (m *OrderedMap) Contains(key int) bool {
	return containsNode(m.root, key)
}

func containsNode(n *node, key int) bool {
	if n == nil {
		return false
	}

	if key < n.key {
		return containsNode(n.left, key)
	}
	if key > n.key {
		return containsNode(n.right, key)
	}

	return true
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	inOrder(m.root, action)
}

func inOrder(n *node, action func(int, int)) {
	if n == nil {
		return
	}

	inOrder(n.left, action)
	action(n.key, n.value)
	inOrder(n.right, action)
}
