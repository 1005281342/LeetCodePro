package main

type CQueue struct {
	List []int
}

func Constructor() CQueue {
	lst := make([]int, 0, 10000)
	cq := new(CQueue)
	cq.List = lst
	return *cq
}

func (this *CQueue) AppendTail(value int) {
	this.List = append(this.List, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.List) == 0 {
		return -1
	}
	ans := this.List[0]
	this.List = this.List[1:]
	return ans
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
