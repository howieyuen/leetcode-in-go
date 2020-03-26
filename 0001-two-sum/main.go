package _001_Two_Sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if val, ok := m[n]; ok {
			return []int{val, i}
		}
		m[target-n] = i
	}
	return nil
}
