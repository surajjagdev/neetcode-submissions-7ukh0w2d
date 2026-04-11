type MinStack struct {
	Stk []int
	MinStk []int
}

func Constructor() MinStack {
	return MinStack{
		Stk: make([]int, 0),
		MinStk: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.Stk = append(this.Stk, val)

	// is this element less than current min
	if len(this.MinStk) == 0 || val <= this.MinStk[len(this.MinStk)-1] {
		this.MinStk = append(this.MinStk, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.Stk) == 0 {
		return
	}

	val := this.Stk[len(this.Stk)-1]
	this.Stk = this.Stk[:len(this.Stk)-1]

	if val == this.MinStk[len(this.MinStk)-1] {
		this.MinStk = this.MinStk[:len(this.MinStk)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Stk) == 0 {
		return -1
	}

	return this.Stk[len(this.Stk)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.MinStk) == 0 {
		return -1
	}

	return this.MinStk[len(this.MinStk)-1]
}
