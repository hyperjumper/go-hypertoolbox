package collection


import (
	"errors"
	"sync"
)

var (
	CollectionFullError  = errors.New("collection is full")
	CollectionEmptyError = errors.New("collection is empty")
)

// Collection collection interface that define how a collection works.
type Collection interface {
	Push(obj interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	PeekAt(index int) (interface{}, error)
	Capacity() int
	IsEmpty() bool
	IsFull() bool
	Length() int
	Elements() []interface{}
	Clear()
}

// NewQueue creates a new queue (FIFO) implementation with maximum capacity.
// Capacity of 0 means that it has no capacity limit.
func NewQueue(cap int) Collection {
	return &Queue{
		capacity: cap,
		elements: make([]interface{}, 0),
		lock:     sync.Mutex{},
	}
}

// Queue is a structure that implement FIFO structure.
type Queue struct {
	capacity int
	elements []interface{}
	lock     sync.Mutex
}

// Clear will clear out the queue, render it empty.
func (q *Queue) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.elements = make([]interface{}, 0)
}

// Elements get the content of this queue as element.
// the first element of the slice is the head of queue where element will be added.
func (q *Queue) Elements() []interface{} {
	return q.elements
}

// Capacity returns the queue's capacity, if its returns 0 means that it has no capacity limit.
func (q *Queue) Capacity() int {
	return q.capacity
}

// IsEmpty Check if a queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// IsFull checks if a queue is full or not
func (q *Queue) IsFull() bool {
	if q.capacity == 0 {
		return false
	}
	return len(q.elements) == q.capacity
}

// Push a data into the queue, the data will be put on the head of the queue,
// returns CollectionFullError if the queue is already full according to the maxLen.
func (q *Queue) Push(obj interface{}) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.capacity > 0 && len(q.elements) >= q.capacity {
		return CollectionFullError
	}

	q.elements = append(q.elements, obj)
	return nil
}

// Pop a data from the queue, The data returned is taken from the tail of the queue. The returned data is taken out
// from the queue.
// returns CollectionEmptyError if the queue is already empty
func (q *Queue) Pop() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.elements) > 0 {
		tail := q.elements[0]
		q.elements = q.elements[1:]
		return tail, nil
	}
	return nil, CollectionEmptyError
}

// Peek a data from the queue tail, The data returned is taken from the tail of the queue. The returned data is not
// taken out from the queue.
// returns CollectionEmptyError if the queue is already empty
// This call is equals to PeekAt(0)
func (q *Queue) Peek() (interface{}, error) {
	return q.PeekAt(0)
}

// PeekAt will peek data from the queue on the specified index, offset-ed from queue tail. The returned data is not
// taken out from the queue
// returns CollectionEmptyError if the queue is already empty
func (q *Queue) PeekAt(index int) (interface{}, error) {
	if len(q.elements) > 0 {
		tail := q.elements[index]
		return tail, nil
	}
	return nil, CollectionEmptyError
}

// Length returns the number of data within the queue
func (q *Queue) Length() int {
	return len(q.elements)
}

// NewStack creates a new stack (FILO) implementation with maximum capacity.
// Capacity of 0 means that it has no capacity limit.
func NewStack(cap int) Collection {
	return &Stack{
		capacity: cap,
		elements: make([]interface{}, 0),
		lock:     sync.Mutex{},
	}
}

// Stack is a structure that implement FILO structure.
type Stack struct {
	capacity int
	elements []interface{}
	lock     sync.Mutex
}

// Elements get the content of this stack as elements.
// the last element of the slice is the head of stack where element will be added.
func (s *Stack) Elements() []interface{} {
	return s.elements
}

// Capacity returns the stack's capacity, if its returns 0 means that it has no capacity limit.
func (s *Stack) Capacity() int {
	return s.capacity
}

// Push a data into the stack, the data will be put on the head of the queue,
// returns CollectionFullError if the stack is already full according to the maxLen.
func (s *Stack) Push(obj interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.capacity > 0 && len(s.elements) >= s.capacity {
		return CollectionFullError
	}
	s.elements = append(s.elements, obj)
	return nil
}

// Pop a data from the stack, The data returned is taken from the head of the stack. The returned data is taken out
// from the stack.
// returns CollectionEmptyError if the queue is already empty
func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.elements) > 0 {
		tail := s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
		return tail, nil
	}
	return nil, CollectionEmptyError
}

// Peek a data from the stack head, The data returned is from the head of the stack. The returned data is not
// taken out from the stack
// returns CollectionEmptyError if the stack is already empty
// This function is equals to PeekAt(0)
func (s *Stack) Peek() (interface{}, error) {
	return s.PeekAt(0)
}

// PeekAt will peek data from the stack on the specified index, offset-ed from stack head. The returned data is not
// taken out from the queue
// returns CollectionEmptyError if the queue is already empty
func (s *Stack) PeekAt(index int) (interface{}, error) {
	if len(s.elements) > 0 {
		tail := s.elements[len(s.elements)-(1+index)]
		return tail, nil
	}
	return nil, CollectionEmptyError
}

// Clear will clear out the stack, render it empty.
func (s *Stack) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elements = make([]interface{}, 0)
}

// IsEmpty Check if a stack is empty or not
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// IsFull checks if a queue is full or not
func (s *Stack) IsFull() bool {
	if s.capacity == 0 {
		return false
	}
	return len(s.elements) == s.capacity
}

// Length returns the number of data within the stack
func (s *Stack) Length() int {
	return len(s.elements)
}

