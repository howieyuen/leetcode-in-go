package _017_Letter_Combinations_of_a_Phone_Number

import (
	`fmt`
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	var strs = []string{"23", "345"}
	for _, str := range strs {
		fmt.Println(letterCombinations(str))
	}
}
