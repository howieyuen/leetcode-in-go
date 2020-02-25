package _040_combination_sum_ii

import (
	`fmt`
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
	candidates = []int{2, 5, 2, 1, 2}
	target = 5
	fmt.Println(combinationSum2(candidates, target))
}
