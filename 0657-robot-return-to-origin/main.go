package _657_robot_return_to_origin

func judgeCircle(moves string) bool {
	i, j := 0, 0
	for _, move := range moves {
		switch move {
		case 'U':
			i--
		case 'D':
			i++
		case 'L':
			j--
		case 'R':
			j++
		}
	}
	return i == 0 && j == 0
}
