package _227_basic_calculator_ii

func calculate(s string) int {
	stack := []int{}
	sign, num := '+', 0
	for i, c := range s {
		n, ok := isDigit(c)
		if ok {
			num = 10*num + n
		}
		if !ok && c != ' ' || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				x := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, x*num)
			case '/':
				x := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, x/num)
			}
			num, sign = 0, c
		}
	}
	sum := 0
	for _, num := range stack {
		sum += num
	}
	return sum
}

func isDigit(c rune) (int, bool) {
	num := int(c - '0')
	if num >= 0 && num <= 9 {
		return num, true
	}
	return 0, false
}
