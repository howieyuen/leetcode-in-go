package _6_26

func calculate(s string) int {
	var nums []int

	var buildNum = func(i int) (int, int) {
		var num = 0
		started := false
		for ; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				started = true
				num = num*10 + int(s[i]-'0')
			} else if s[i] == ' ' && !started {
				continue
			} else {
				break
			}
		}
		return i, num
	}

	for i := 0; i < len(s); {
		switch s[i] {
		case '+':
			i++
		case '-':
			j, num := buildNum(i + 1)
			nums = append(nums, -num)
			i = j
		case '*':
			j, num := buildNum(i + 1)
			nums[len(nums)-1] = nums[len(nums)-1] * num
			i = j
		case '/':
			j, num := buildNum(i + 1)
			nums[len(nums)-1] = nums[len(nums)-1] / num
			i = j
		case ' ':
			i++
		default:
			j, num := buildNum(i)
			nums = append(nums, num)
			i = j
		}
	}

	var ans int
	for i := range nums {
		ans += nums[i]
	}
	return ans
}
