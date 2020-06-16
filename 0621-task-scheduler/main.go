package _621_task_scheduler

import (
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	var times = make([]int, 26)
	for _, t := range tasks {
		times[int(t-'A')]++
	}
	sort.Ints(times)
	maxTime := times[25] - 1
	idle := maxTime * n
	for i := 24; i >= 0 && times[i] > 0; i-- {
		idle -= min(times[i], maxTime)
	}
	if idle > 0 {
		return idle + len(tasks)
	}
	return len(tasks)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
