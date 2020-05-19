package _0

import "math"


/*
see https://leetcode-cn.com/problems/min-stack/
*/
type MinStack struct {
	data []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  math.MaxInt64,
	}
}

func (m *MinStack) Push(x int) {
	m.data = append(m.data, x)
	if x < m.min {
		m.min = x
	}
}

func (m *MinStack) Pop() {
	if len(m.data) > 0 {
		cur := m.data[len(m.data)-1]
		m.data = m.data[:len(m.data)-1]
		if cur == m.min {
			m.min = math.MaxInt64
			for _, x := range m.data {
				if x < m.min {
					m.min = x
				}
			}
		}
	}
}

func (m *MinStack) Top() int {
	if len(m.data) == 0 {
		return math.MaxInt64
	}
	return m.data[len(m.data)-1]
}

func (m *MinStack) Min() int {
	return m.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
