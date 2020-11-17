package _030_matrix_cells_in_distance_order

import (
	`sort`
)

type node struct {
	location []int
	distance int
}

type nodes []node

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	nodes := make(nodes, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			nodes = append(nodes, node{[]int{i, j}, getManhattan(r0, c0, i, j)})
		}
	}
	sort.Sort(nodes)
	result := make([][]int, 0)
	for i := 0; i < nodes.Len(); i++ {
		result = append(result, nodes[i].location)
	}
	return result
}

func getManhattan(r1, c1, r2, c2 int) int {
	return abs(r1-r2) + abs(c1-c2)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Less(i, j int) bool {
	return n[i].distance < n[j].distance
}

func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func allCellsDistOrder1(R int, C int, r0 int, c0 int) [][]int {
	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}
	queue := make([][]int, 0)
	output := make([][]int, 0)
	queue = append(queue, []int{r0, c0})
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		r := point[0]
		c := point[1]
		if visited[r][c] {
			continue
		}
		visited[r][c] = true
		output = append(output, point)
		if r-1 >= 0 {
			queue = append(queue, []int{r - 1, c})
		}
		if r+1 < R {
			queue = append(queue, []int{r + 1, c})
		}
		if c-1 >= 0 {
			queue = append(queue, []int{r, c - 1})
		}
		if c+1 < C {
			queue = append(queue, []int{r, c + 1})
		}
	}
	return output
}
