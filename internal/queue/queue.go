package queue

type (
	//Queue 队列
	Queue struct {
		top    *node
		rear   *node
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
func New() *Queue {
	return &Queue{nil, nil, 0}
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
	n := &node{nil, nil, v}
	if q.length == 0 {
		q.top = n
		q.rear = q.top
	} else {
		n.pre = q.rear
		q.rear.next = n
		q.rear = n
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
	q.length--
	return n.value
}
