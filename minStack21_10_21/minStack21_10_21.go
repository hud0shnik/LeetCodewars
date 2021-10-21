package main

type MinStack struct {
	stack []int
	min   []int
	size  int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
		size:  0,
	}
}

func (this *MinStack) Push(num int) {
	this.stack = append(this.stack, num)
	min := num
	if this.size > 0 && this.min[this.size-1] < min {
		min = this.min[this.size-1]
	}

	this.min = append(this.min, min)
	this.size++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.size-1]
	this.min = this.min[:this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.size-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
