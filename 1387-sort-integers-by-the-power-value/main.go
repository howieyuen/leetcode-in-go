package _387_sort_integers_by_the_power_value

import "sort"

func getKth(lo int, hi int, k int) int {
	var dict = make(map[int]int)
	var weights func(x int) int
	weights = func(x int) int {
		if _, ok := dict[x]; !ok {
			if x == 1 {
				dict[x] = 0
			} else if x%2 == 0 {
				dict[x] = weights(x>>1) + 1
			} else {
				dict[x] = weights(x*3+1) + 1
			}
		}
		return dict[x]
	}
	var arr = make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		arr[i-lo] = i
	}
	sort.Slice(arr, func(i, j int) bool {
		x, y := weights(arr[i]), weights(arr[j])
		if x != y {
			return x < y
		}
		return arr[i] < arr[j]
	})
	return arr[k-1]
}

func getKth1(lo, hi, k int) int {
	arr := make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		arr[i-lo] = i
	}
	var weights = func() (weights []int) {
		weights = make([]int, hi+1)
		for i := range weights {
			now := i
			weights[i] = 0
			for now > 1 {
				if now < i {
					weights[i] += weights[now]
					break
				}
				if now%2 == 0 {
					weights[i]++
					now >>= 1
				} else {
					weights[i]++
					now = now*3 + 1
				}
			}
		}
		return weights
	}()

	sort.Slice(arr, func(i, j int) bool {
		return weights[arr[i]] < weights[arr[j]] || weights[arr[i]] == weights[arr[j]] && arr[i] < arr[j]
	})
	return arr[k-1]
}

func getKth2(lo, hi, k int) int {
	arr := make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		arr[i-lo] = i
	}
	sort.Sort(byWeight(arr))
	return arr[k-1]
}

var weight []int

func init() {
	weight = make([]int, 1001)
	for i := range weight {
		now := i
		weight[i] = 0
		for now > 1 {
			if now < i {
				weight[i] += weight[now]
				break
			}
			if now%2 == 0 {
				weight[i]++
				now >>= 1
			} else {
				weight[i]++
				now = now*3 + 1
			}
		}
	}
}

type byWeight []int

func (w byWeight) Len() int {
	return len(w)
}

func (w byWeight) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w byWeight) Less(i, j int) bool {
	return weight[w[i]] < weight[w[j]] || weight[w[i]] == weight[w[j]] && w[i] < w[j]
}
