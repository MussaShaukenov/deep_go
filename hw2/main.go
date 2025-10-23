package main

import (
	"fmt"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	front  int
	rear   int
}

// Создать очередь с определенным размером буффера
func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		front:  -1,
		rear:   -1,
	}
}

// добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	if q.front == -1 { // empty queue
		q.front = 0
		q.rear = 0
	} else {
		q.rear = (q.rear + 1) % len(q.values)
	}

	q.values[q.rear] = value
	return true
}

// удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	// element = q.values[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % len(q.values)
	}

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

	return q.values[q.rear]
}

// проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return q.front == -1
}

// проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	if q.front == 0 && q.rear == len(q.values)-1 {
		return true
	}

	if q.front == q.rear+1 {
		return true
	}

	return false
}

func main() {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	fmt.Println(queue.Empty())
}
