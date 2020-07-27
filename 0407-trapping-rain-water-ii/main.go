package _407_trapping_rain_water_ii

import (
	"container/heap"
)

// 1. 最外围作为围栏入最小堆；
// 2. 出堆，当前值是围栏的最矮的一个，搜索最矮的围栏周围;
// 3. 内部柱子有比最矮围栏还矮的，则可以灌水；
//    更新此柱子高度（选择原来高度和最矮围栏高度最大的那一个），并将这个柱子作为新的一格围栏和原来围栏组成新的外围栏，入堆
// 4. 重复 2，3;
func trapRainWater(heightMap [][]int) int {
	rows := len(heightMap)
	if rows == 0 {
		return 0
	}
	cols := len(heightMap[0])
	if cols == 0 {
		return 0
	}
	var visited = make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	var pq PriorityQueue
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 最外围围栏入堆
			if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
				heap.Push(&pq, &Item{i: i, j: j, height: heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}
	var ans int
	var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		for _, dir := range directions {
			i, j := item.i+dir[0], item.j+dir[1]
			if i < 0 || i >= rows || j < 0 || j >= cols || visited[i][j] {
				continue
			}
			// 围栏比内部高，可以灌水
			if item.height > heightMap[i][j] {
				ans += item.height - heightMap[i][j]
			}
			// 忽略当前围栏，在(i, j)处新建围栏
			visited[i][j] = true
			// 新的围栏入堆
			heap.Push(&pq, &Item{i: i, j: j, height: max(heightMap[i][j], item.height)})
		}
	}
	return ans
}

type Item struct {
	i, j   int
	height int
}

type PriorityQueue []*Item

var _ heap.Interface = &PriorityQueue{}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].height < pq[j].height
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
