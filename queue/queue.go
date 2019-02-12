package queue

//An FIFO queue.
type Queue []interface{}

//Pushes the element into the queue.
// 	e.g. go.Push(123)
func (q *Queue) Push(v interface{}) {
	//在內部確認類型為int 若非int則會報錯
	*q = append(*q, v.(int))
}

//Pops eleement from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

//Returns if th queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
