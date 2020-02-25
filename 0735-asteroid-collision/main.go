package _735_asteroid_collision

func asteroidCollision(asteroids []int) []int {
	var ans []int
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			ans = append(ans, asteroids[i])
		} else {
			if len(ans) == 0 || ans[len(ans)-1] < 0 {
				ans = append(ans, asteroids[i])
			} else if -asteroids[i] > ans[len(ans)-1] {
				ans = ans[:len(ans)-1]
				i--
			} else if -asteroids[i] == ans[len(ans)-1] {
				ans = ans[:len(ans)-1]
			}
		}
	}
	return ans
}
