package _498_diagonal_traverse

/*
每层的索引和相等:
1. 假设矩阵无限大；
2. 索引和为{偶}数，向上遍历，{横}索引值递减，遍历值依次是(x,0),(x-1,1),(x-2,2),...,(0,x)
3. 索引和为{奇}数，向下遍历，{纵}索引值递减，遍历值依次是(0,y),(1,y-1),(2,y-2),...,(y,0)

每层的索引和:
0:              (00)
1:            (01)(10)
2:          (20)(11)(02)
3:        (03)(12)(21)(30)
4:      (40)(31)(22)(13)(04)
5:    (05)(14)(23)(32)(41)(50)
6:  (60)(51)................(06)

按照“层次”遍历，依次append在索引边界内的值即可
*/
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	var nums = make([]int, 0, m*n)
	for x := 0; x < m+n; x++ {
		if x%2 == 0 {
			for i := x; i >= 0; i-- {
				j := x - i
				if i < m && j < n {
					nums = append(nums, matrix[i][j])
				} else if j >= n {
					break
				} else {
					continue
				}
			}
		} else {
			for j := x; j >= 0; j-- {
				i := x - j
				if i < m && j < n {
					nums = append(nums, matrix[i][j])
				} else if i >= m {
					break
				} else {
					continue
				}
			}
		}
	}
	return nums
}

func findDiagonalOrder1(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, cols := len(matrix), len(matrix[0])
	var nums = make([]int, 0, rows*cols)
	flag := true
	for i := 0; i < rows+cols; i++ {
		var pm, pn = rows, cols
		if !flag {
			pm, pn = cols, rows
		}
		var x int
		if i < pm {
			x = i
		} else {
			x = pm - 1
		}
		var y = i - x
		for x >= 0 && y < pn {
			if flag {
				nums = append(nums, matrix[x][y])
			} else {
				nums = append(nums, matrix[y][x])
			}
			x--
			y++
		}
		flag = !flag
	}
	return nums
}
