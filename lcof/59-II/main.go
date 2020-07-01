package _9_II

type MaxQueue struct {
	data []int
	max  []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		data: nil,
		max:  nil,
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)
	for i, v := range this.max {
		if v < value {
			this.max = this.max[:i]
			break
		}
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) == 0 {
		return -1
	}
	front := this.data[0]
	this.data = this.data[1:]
	if front == this.max[0] {
		this.max = this.max[1:]
	}
	return front
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
