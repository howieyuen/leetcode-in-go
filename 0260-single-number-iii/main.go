package _260_single_number_iii

func singleNumber(nums []int) []int {
	bitMask := 0
	for _, num := range nums {
		bitMask ^= num
	}
	// 保留bisMask中最右边的1
	// 这个1就是只出现1次的数的其中一个
	diff := bitMask & (-bitMask)
	x := 0
	for _, num := range nums {
		// num存在最右边的1，可能是出现一次的数
		// 去除出现2次的数
		if num&diff != 0 {
			x ^= num
		}
	}
	return []int{x, bitMask ^ x}
}
