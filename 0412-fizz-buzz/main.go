package _412_fizz_buzz

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	var ans = make([]string, n)
	for i := range ans {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			ans[i] = "FizzBuzz"
		} else if (i+1)%3 == 0 {
			ans[i] = "Fizz"
		} else if (i+1)%5 == 0 {
			ans[i] = "Buzz"
		} else {
			ans[i] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
