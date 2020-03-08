package _860_lemonade_change

/*
	找零 5 10 20
*/
func lemonadeChange(bills []int) bool {
	if len(bills) == 0 {
		return true
	}
	five := 0
	ten := 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else if bill == 20 {
			if ten >= 1 && five >= 1 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
