package _365_water_and_jug_problem

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	return z%gcd(x, y) == 0
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for x%y != 0 {
		x, y = y, x%y
	}
	return x
}
