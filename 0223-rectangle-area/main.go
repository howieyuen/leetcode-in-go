package _223_rectangle_area

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area1 := innerArea(A, B, C, D)
	area2 := innerArea(E, F, G, H)
	if A >= G || B >= H || C <= E || D <= F {
		return area1 + area2
	}
	x1 := max(A, E)
	y1 := max(B, F)
	x2 := min(C, G)
	y2 := min(D, H)
	return area1 - abs(x1-x2)*abs(y1-y2) + area2
}

func innerArea(a, b, c, d int) int {
	return abs(a-c) * abs(b-d)
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
