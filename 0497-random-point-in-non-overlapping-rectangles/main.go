package _497_random_point_in_non_overlapping_rectangles

import (
	"math/rand"
)

type Solution struct {
	// 矩阵集合
	rects [][]int
	// 矩阵中坐标点的前缀和
	points []int
	// 坐标点总数
	total int
}

func Constructor(rects [][]int) Solution {
	n := len(rects)
	var points = make([]int, n+1)
	for i := range rects {
		x1, y1, x2, y2 := rects[i][0], rects[i][1], rects[i][2], rects[i][3]
		points[i+1] = points[i] + (y2-y1+1)*(x2-x1+1)
	}
	return Solution{
		rects:  rects,
		points: points,
		total:  points[n],
	}
}

func (s *Solution) Pick() []int {
	// 随机第i个坐标点
	i := rand.Intn(s.total)
	l, r := 1, len(s.points)-1
	for l < r {
		m := (l + r) / 2
		if s.points[m] <= i {
			l = m + 1
		} else {
			r = m
		}
	}
	// 第i个坐标点是第l个矩阵中的第j个坐标
	j := i - s.points[l-1]
	// 第l个矩阵的坐标(x1, y1), (x2, y2)
	x1, y1, y2 := s.rects[l-1][0], s.rects[l-1][1], s.rects[l-1][3]
	height := y2 - y1 + 1
	return []int{j/height + x1, j%height + y1}
}
