package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	//在內部確認類型為int 若非int則會報錯
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
