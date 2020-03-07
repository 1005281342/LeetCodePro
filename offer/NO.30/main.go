package main

type MinStack struct {
	List   []int
	Length int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	d := make([]int, 0, 20000)
	return MinStack{List: d}
}

func (this *MinStack) Push(x int) {
	this.List = append(this.List, x)
	this.Length += 1
}

func (this *MinStack) Pop() {
	this.Length -= 1
	this.List = this.List[:this.Length]
}

func (this *MinStack) Top() int {
	return this.List[this.Length-1]
}

func (this *MinStack) Min() int {
	ans := (2 << 31) - 1
	for _, v := range this.List {
		if v < ans {
			ans = v
		}
	}
	return ans
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
