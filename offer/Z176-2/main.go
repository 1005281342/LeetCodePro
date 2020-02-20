package main

type ProductOfNumbers struct {
	List   []int
	Length int
}

func Constructor() ProductOfNumbers {
	lst := make([]int, 0, 101)
	p := ProductOfNumbers{List: lst}
	return p
}

func (this *ProductOfNumbers) Add(num int) {
	this.List = append(this.List, num)
	this.Length += 1
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	ans := 1
	for i := this.Length - k; i < this.Length; i++ {
		ans *= this.List[i]
	}
	return ans
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
