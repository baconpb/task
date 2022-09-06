package queue

func New(list ...interface{}) *Queue {
	return &Queue{list: list}
}

type Queue struct {
	list []interface{}
}

func (q *Queue) Push(v interface{}) {
	q.list = append(q.list, v)
}

func (q *Queue) Pop() (interface{}, bool) {
	if len(q.list) == 0 {
		var t interface{}
		return t, false
	}
	var v = q.list[0]
	q.list = q.list[1:]
	return v, true
}

func (q *Queue) Top() (interface{}, bool) {
	if len(q.list) == 0 {
		var t interface{}
		return t, false
	}
	return q.list[0], true
}

func (q *Queue) Size() int {
	return len(q.list)
}
