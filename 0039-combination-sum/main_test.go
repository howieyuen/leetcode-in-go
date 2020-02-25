package _039_combination_sum

import (
	`fmt`
	"testing"
)

func Test_combinationSum(t *testing.T) {
	var candidates = []int{2, 3, 6, 7}
	var target = 7
	fmt.Println(combinationSum(candidates, target))
	candidates = []int{2, 3, 5}
	target = 8
	fmt.Println(combinationSum(candidates, target))
}
