package queue

import "github.com/xiecat/fofax/internal/printer"

type (
	//Queue 队列
	Queue struct {
		top    *node
		rear   *node
		Uniq   bool
		umap   map[string]bool
		length int
	}
	//双向链表节点
	node struct {
		pre   *node
		next  *node
		value string
	}
)

// New Create a new queue
func New(isUniq bool) *Queue {
	return &Queue{
		top:    nil,
		rear:   nil,
		Uniq:   isUniq,
		umap:   make(map[string]bool),
		length: 0,
	}
}

// Len 获取队列长度
func (q *Queue) Len() int {
	return q.length
}

// Any 返回true队列不为空
func (q *Queue) Any() bool {
	return q.length > 0
}

// Peek 返回队列顶端元素
func (q *Queue) Peek() string {
	if q.top == nil {
		return ""
	}
	return q.top.value
}

// Push 入队操作
func (q *Queue) Push(v string) {

	if _, ok := q.umap[v]; ok && q.Uniq {
		printer.Infof("The %s to be queried is repeated", v)
		return
	}

	n := &node{nil, nil, v}
	if q.length == 0 {
		q.top = n
		q.rear = q.top
	} else {
		n.pre = q.rear
		q.rear.next = n
		q.rear = n
	}
	if q.Uniq {
		q.umap[v] = true
	}
	q.length++
}

// Pop 出队操作
func (q *Queue) Pop() string {
	if q.length == 0 {
		return ""
	}
	n := q.top
	if q.top.next == nil {
		q.top = nil
	} else {
		q.top = q.top.next
		q.top.pre.next = nil
		q.top.pre = nil
	}

	if q.Uniq {
		delete(q.umap, n.value)
	}
	q.length--
	return n.value
}
