package main

type MyStack struct {
	Lst    []int
	Length int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	v := make([]int, 0, 1024)
	return MyStack{Lst: v}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Lst = append(this.Lst, x)
	this.Length += 1
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {

	this.Length -= 1
	v := this.Lst[this.Length]
	this.Lst = this.Lst[:this.Length]
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Lst[this.Length-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Length <= 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
