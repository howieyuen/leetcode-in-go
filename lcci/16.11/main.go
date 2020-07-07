package _6_11

// 0<=i<=j<=k
// shorter*(k−i)+longer*i
// shorter*(k−j)+longer*j
// (shorter*(k−i)+longer*i) - (shorter*(k−j)+longer*j) = (longer−shorter)*(i−j)
// 显然大于0，因为k+1种组合方式，长度各不相同
func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	lengths := make([]int, k+1)
	for i := 0; i <= k; i++ {
		lengths[i] = shorter*(k-i) + longer*i
	}
	return lengths
}
