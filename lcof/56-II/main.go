package _6_II

func singleNumber(nums []int) int {
	var a, b int
	for i := range nums {
		a = a ^ nums[i]&^b
		b = b ^ nums[i]&^a
	}
	return a
}

// 把所有的数转化为3进制数，不进位求和
// 和的每一位，如果是3的倍数，表示要找的数在这一位是0，反之是1
func singleNumber1(nums []int) int {
	var bits = make([]int, 32)
	for i := range nums {
		for j := 0; j < 32; j++ {
			bits[j] += nums[i] & 1
			nums[i] >>= 1
		}
	}
	res, base := 0, 3
	for i := 0; i < 32; i++ {
		res <<= 1
		res |= (bits[31-i] % base)
	}
	return res
}
