package main

import "fmt"

func shoppingOffers(price []int, special [][]int, needs []int) int {
	var shopping func(needs []int) int
	shopping = func(needs []int) int {
		res := dot(price, needs)
		for _, packet := range special {
			tmp := make([]int, len(needs))
			copy(tmp, needs)
			var i int
			for i = 0; i < len(needs); i++ {
				diff := tmp[i] - packet[i]
				if diff < 0 {
					break
				}
				tmp[i] = diff
			}
			if i == len(needs) {
				res = min(res, packet[i]+shopping(tmp))
			}
		}
		return res
	}
	return shopping(needs)
}

func shoppingOffers1(price []int, special [][]int, needs []int) int {
	var memory = make(map[string]int)
	var shopping func(needs []int) int
	shopping = func(needs []int) int {
		if v, ok := memory[fmt.Sprintf("%v", needs)]; ok {
			return v
		}
		res := dot(price, needs)
		for _, packet := range special {
			tmp := make([]int, len(needs))
			copy(tmp, needs)
			var i int
			for i = 0; i < len(needs); i++ {
				left := tmp[i] - packet[i]
				if left < 0 {
					break
				}
				tmp[i] = left
			}
			if i == len(needs) {
				res = min(res, packet[i]+shopping(tmp))
			}
		}
		memory[fmt.Sprintf("%v", needs)] = res
		return res
	}
	return shopping(needs)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dot(price []int, needs []int) int {
	var res int
	for i := range needs {
		res += needs[i] * price[i]
	}
	return res
}
