package _260_single_number_iii

func singleNumber(nums []int) []int {
	bitMask := 0
	for _, num := range nums {
		bitMask ^= num
	}
	diff := bitMask ^ (-bitMask)
	x := 0
	for _, num := range nums {
		if num&diff != 0 {
			x ^= num
		}
	}
	return []int{x, bitMask ^ x}
}
