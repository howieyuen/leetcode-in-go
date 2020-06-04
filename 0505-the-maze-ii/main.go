package _505_the_maze_ii

import (
	"container/heap"
	"math"
)

func shortestDistance1(maze [][]int, start []int, destination []int) int {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return -1
	}
	
	rows, cols := len(maze), len(maze[0])
	sx, sy := start[0], start[1]
	if sx < 0 || sx >= rows || sy < 0 || sy >= cols {
		return -1
	}
	ex, ey := destination[0], destination[1]
	if ex < 0 || ex >= rows || ey < 0 || ey >= cols {
		return -1
	}
	
	if maze[sx][sy] == 1 || maze[ex][ey] == 1 {
		return -1
	}
	
	var dis = make([][]int, rows)
	for i := range dis {
		dis[i] = make([]int, cols)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt64
		}
	}
	dis[sx][sy] = 0
	
	var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		for _, dir := range directions {
			a := i + dir[0]
			b := j + dir[1]
			count := 0
			for a >= 0 && a < rows && b >= 0 && b < cols && maze[a][b] == 0 {
				a += dir[0]
				b += dir[1]
				count++
			}
			if dis[i][j]+count < dis[a-dir[0]][b-dir[1]] {
				dis[a-dir[0]][b-dir[1]] = dis[i][j] + count
				dfs(a-dir[0], b-dir[1])
			}
		}
	}
	dfs(sx, sy)
	if dis[ex][ey] == math.MaxInt64 {
		return -1
	}
	return dis[ex][ey]
}

// An Item is something we manage in a distance queue.
type Item struct {
	x        int // The x of the item; arbitrary.
	y        int // The y of the item in the heap.
	distance int // The distance of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, distance so we use greater than here.
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[:n-1]
	return item
}

func shortestDistance(maze [][]int, start []int, destination []int) int {
	rows, cols := len(maze), len(maze[0])
	var dis = make([][]int, rows)
	for i := range dis {
		dis[i] = make([]int, cols)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt64
		}
	}
	dis[start[0]][start[1]] = 0
	
	var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		x:        start[0],
		y:        start[1],
		distance: 0,
	}
	heap.Init(&pq)
	
	for pq.Len() > 0 {
		item := pq.Pop().(*Item)
		if dis[item.x][item.y] < item.distance {
			continue
		}
		for _, dir := range directions {
			a := item.x + dir[0]
			b := item.y + dir[1]
			count := 0
			for a >= 0 && a < rows && b >= 0 && b < cols && maze[a][b] == 0 {
				a += dir[0]
				b += dir[1]
				count++
			}
			if dis[item.x][item.y]+count < dis[a-dir[0]][b-dir[1]] {
				dis[a-dir[0]][b-dir[1]] = dis[item.x][item.y] + count
				pq.Push(&Item{
					x:        a - dir[0],
					y:        b - dir[1],
					distance: dis[a-dir[0]][b-dir[1]],
				})
			}
		}
	}
	if dis[destination[0]][destination[1]] == math.MaxInt64 {
		return -1
	}
	return dis[destination[0]][destination[1]]
}
