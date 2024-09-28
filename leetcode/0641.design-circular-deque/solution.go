package main

type DequeNode struct {
	val  int
	next *DequeNode
	prev *DequeNode
}

type CircularDeque struct {
	head     *DequeNode
	tail     *DequeNode
	size     int
	capacity int
}

func NewCircularDeque(capacity int) *CircularDeque {
	return &CircularDeque{nil, nil, 0, capacity}
}

func (c *CircularDeque) InsertFront(val int) bool {
	if c.IsFull() {
		return false
	}

	newHead := &DequeNode{val: val, prev: nil, next: nil}

	if c.head != nil {
		newHead.next = c.head
		c.head.prev = newHead
	} else {
		c.tail = newHead
	}

	c.head = newHead
	c.size++

	return true
}

func (c *CircularDeque) InsertLast(val int) bool {
	if c.IsFull() {
		return false
	}

	newTail := &DequeNode{val: val, prev: nil, next: nil}

	if c.tail != nil {
		newTail.prev = c.tail
		c.tail.next = newTail
	} else {
		c.head = newTail
	}

	c.tail = newTail
	c.size++

	return true
}

func (c *CircularDeque) DeleteFront(val int) bool {
	if c.IsEmpty() {
		return false
	}

	if c.head == c.tail {
		c.head = nil
		c.tail = nil
	} else {
		oldHead := c.head
		c.head = oldHead.next
		c.head.prev = nil
		oldHead = nil
	}

	c.size--

	return true
}

func (c *CircularDeque) DeleteLast(val int) bool {
	if c.IsEmpty() {
		return false
	}

	if c.head == c.tail {
		c.head = nil
		c.tail = nil
	} else {
		oldTail := c.tail
		c.tail = oldTail.prev
		c.tail.next = nil
		oldTail = nil
	}

	c.size--

	return true
}

func (c *CircularDeque) GetFront(val int) int {
	if c.IsEmpty() {
		return -1
	}

	return c.head.val
}

func (c *CircularDeque) GetLast(val int) int {
	if c.IsEmpty() {
		return -1
	}

	return c.tail.val
}

func (c *CircularDeque) IsEmpty() bool {
	return c.size == 0
}

func (c *CircularDeque) IsFull() bool {
	return c.size == c.capacity
}
