package queue

import (
	"sync"
)

func NewBlock() *Block {
	var queue = &Block{}
	queue.cond = sync.NewCond(new(sync.RWMutex))
	queue.status = Status{Wait: 0, Len: 0}
	return queue
}

type Status struct {
	Wait int
	Len  int
}

type Block struct {
	cond    *sync.Cond
	storage []interface{}
	status  Status
}

func (queue *Block) Push(v interface{}) {

	queue.cond.L.Lock()

	queue.storage = append(queue.storage, v)
	queue.status.Len++

	if queue.status.Wait > 0 {
		queue.cond.Signal()
	}

	queue.cond.L.Unlock()
}

func (queue *Block) Pop() interface{} {

	queue.cond.L.Lock()

	queue.status.Wait++

	for {
		if len(queue.storage) > 0 {
			var r = queue.storage[0]
			queue.storage = queue.storage[1:]
			queue.status.Wait--
			queue.status.Len--
			queue.cond.L.Unlock()
			return r
		}
		queue.cond.Wait()
	}
}

func (queue *Block) Status() Status {
	return queue.status
}
