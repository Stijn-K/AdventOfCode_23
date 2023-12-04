package common

import "log"

type Queue struct {
	items []interface{}
}

func (q *Queue) Push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Pop() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	var item interface{}
	item, q.items = q.items[0], q.items[1:]
	return item
}

func (q *Queue) Peek() interface{} {
	if len(q.items) > 0 {
		return q.items[0]
	}
	return nil
}

func (q *Queue) Print() {
	for i, item := range q.items {
		log.Printf("%d: %v", i, item)
	}
}
