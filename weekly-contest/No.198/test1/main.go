package test1

func numWaterBottles(numBottles int, numExchange int) int {
	var ans = numBottles
	for numBottles >= numExchange {
		tmp := numBottles / numExchange
		remainder := numBottles % numExchange
		ans += tmp
		numBottles = tmp + remainder
	}
	return ans
}
