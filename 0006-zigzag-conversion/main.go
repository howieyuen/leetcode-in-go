package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	goingDown := false
	rowStrs := make([][]byte, min(len(s), numRows))
	for i := range rowStrs {
		rowStrs[i] = []byte{}
	}
	curRow := 0
	for i := range s {
		rowStrs[curRow] = append(rowStrs[curRow], s[i])
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		} else {
			curRow += -1
		}
	}
	for i := 1; i < len(rowStrs); i++ {
		rowStrs[0] = append(rowStrs[0], rowStrs[i]...)
	}
	return string(rowStrs[0])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
