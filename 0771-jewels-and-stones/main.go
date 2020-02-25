package _771_jewels_and_stones

func numJewelsInStones(J string, S string) int {
	var maps = make(map[byte]int)
	for i := range J {
		if _, ok := maps[J[i]]; ok {
			maps[J[i]]++
		} else {
			maps[J[i]] = 1
		}
	}
	sum := 0
	for i := range S {
		if n, ok := maps[S[i]]; ok {
			sum += n
		}
	}
	return sum
}
