package _179_largest_number

import (
	`sort`
	`strconv`
	`strings`
)

func largestNumber(nums []int) string {
	var strs Numbers
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	sort.Sort(strs)
	ans := strings.Join(strs, "")
	if ans[0] == '0' {
		ans = "0"
	}
	return ans
}

type Numbers []string

func (n Numbers) Len() int {
	return len(n)
}

func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Numbers) Less(i, j int) bool {
	return n[i]+n[j] > n[j]+n[i]
}
