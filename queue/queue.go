package queue

import "sync"

var mux = sync.Mutex{}

type Queue struct {
	elements []string
}

// Put put data to the end of Queue
func (q *Queue) Put(data string) {
	mux.Lock()
	q.elements = append(q.elements, data)
	mux.Unlock()
}

// Pop pop data from the head of Queue
func (q *Queue) Pop() (string, bool) {
	mux.Lock()
	defer mux.Unlock()
	var val string
	if len(q.elements) == 0 {
		return val, true
	}
	val = q.elements[0]
	q.elements = q.elements[1:]
	return val, len(q.elements) == 0
}

// Size get the size of Queue
func (q *Queue) Size() int {
	return len(q.elements)
}
