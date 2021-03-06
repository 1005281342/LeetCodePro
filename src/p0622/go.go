package main

type MyCircularQueue struct {
	rq *RingQueueBlock
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{rq: NewRingQueueBlock(k)}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	var b = this.rq.IsFull()
	this.rq.Insert(value)
	return !b
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.rq.Empty() {
		return false
	}
	this.rq.Pop()
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.rq.Empty() {
		return -1
	}
	return this.rq.Head().(int)
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.rq.Empty() {
		return -1
	}
	return this.rq.Tail().(int)
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.rq.Empty()
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.rq.IsFull()
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

type RingQueueBlock struct {
	front  int           // 指向队列头部第1个有效数据的位置
	rear   int           // 指向队列尾部（即最后1个有效数据）的下一个位置，即下一个从队尾入队元素的位置
	length int           // 队列长度，非容量，slice动态库容
	nums   []interface{} // 队列元素
}

// NewRingQueueBlock Initialize the structure. Set the size of the queue to be k.
func NewRingQueueBlock(k int) *RingQueueBlock {
	length := k + 1
	return &RingQueueBlock{
		length: length,
		nums:   make([]interface{}, length),
	}
}

func (q *RingQueueBlock) Len() int {
	if q.front > q.rear {
		return q.rear - q.front + q.length
	}
	return q.rear - q.front
}

func (q *RingQueueBlock) Head() interface{} {
	if q.Empty() {
		return nil
	}

	return q.nums[q.front]
}

func (q *RingQueueBlock) Tail() interface{} {
	if q.Empty() {
		return nil
	}
	var pos int // 其实就是rear-1 //pos = (q.rear - 1 + q.length) % q.length
	if q.rear > 0 {
		pos = q.rear - 1
	} else if q.rear == 0 {
		pos = q.length - 1
	}
	return q.nums[pos]
}

func (q *RingQueueBlock) Insert(x interface{}) interface{} {
	if q.IsFull() {
		return false
	}
	q.nums[q.rear] = x
	q.rear++ // q.rear = (q.rear + 1) % q.length // rear指针后移一位
	if q.rear == q.length {
		q.rear = 0
	}
	return true
}

func (q *RingQueueBlock) Pop() interface{} {
	if q.Empty() {
		return false
	}

	if q.front == q.length-1 { // front指针后移一位 // q.front = (q.front + 1) % q.length
		q.front = 0
	} else {
		q.front += 1
	}
	return true
}

// IsEmpty check the queue is empty.
func (q *RingQueueBlock) Empty() bool {
	return q.front == q.rear
}

// IsFull check the queue is full.
func (q *RingQueueBlock) IsFull() bool {
	return (q.rear+1)%q.length == q.front
}
