package _476_subrectangle_queries

type SubrectangleQueries struct {
	rectangle [][]int
	operators [][5]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		rectangle: rectangle,
	}
}

func (q *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	q.operators = append(q.operators, [5]int{row1, col1, row2, col2, newValue})
}

func (q *SubrectangleQueries) GetValue(row int, col int) int {
	for i := len(q.operators) - 1; i >= 0; i-- {
		if q.operators[i][0] <= row && row <= q.operators[i][2] &&
			q.operators[i][1] <= col && col <= q.operators[i][3] {
			return q.operators[i][4]
		}
	}
	return q.rectangle[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */

/*
type SubrectangleQueries struct {
	rows, cols int
	rectangle  [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		rows:      len(rectangle),
		cols:      len(rectangle[0]),
		rectangle: rectangle,
	}
}

func (q *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			q.rectangle[i][j] = newValue
		}
	}
}

func (q *SubrectangleQueries) GetValue(row int, col int) int {
	return q.rectangle[row][col]
}

*/
