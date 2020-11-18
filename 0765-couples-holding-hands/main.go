package _765_couples_holding_hands

func minSwapsCouples(row []int) int {
	var findMatch = func(n int) int {
		if n%2 == 0 {
			return n + 1
		}
		return n - 1
	}
	
	c := 0
	for i := 0; i < len(row); i += 2 {
		p1 := row[i]
		p2 := findMatch(p1)
		if row[i+1] != p2 {
			var j int
			for ; j < len(row); j++ {
				if row[j] == p2 {
					break
				}
			}
			row[i+1], row[j] = row[j], row[i+1]
			c += 1
		}
	}
	return c
}
