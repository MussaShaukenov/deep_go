package main

import (
	"fmt"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	front  int
	rear   int
	count  int
}

// Создать очередь с определенным размером буффера
func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
	}
}

// добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	q.values[q.rear] = value
	q.rear = (q.rear + 1) % len(q.values)
	q.count++

	return true
}

// удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	q.front = (q.front + 1) % len(q.values)
	q.count--

	return true
}

// получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.front]
}

// получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	last := (q.rear - 1 + len(q.values)) % len(q.values)
	return q.values[last]
}

// проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return q.count == 0
}

// проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	return q.count == len(q.values)
}

func main() {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	fmt.Println(queue.Empty())
}
