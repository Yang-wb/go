package queue

// An FIFO queue
type Queue []interface{}

// Pushes the element into the queue.
//   e.g. q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

//Pops element form head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	//强制转换为int
	return head.(int)
}

//Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
