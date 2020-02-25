package _752_open_the_lock

/*
	打开转盘锁，转化成图算法，每衍生一轮变化，相当于多了一层子节点
	相当于 dfs 寻找最短路径
	对于锁盘上每个数字，依次转动，结果放入队列，作为下次查找的节点；
	每轮遍历队列中所有节点，直到路径匹配
*/
func openLock(deadends []string, target string) int {
	var visited = make(map[string]bool)
	for _, deadend := range deadends {
		visited[deadend] = true
	}
	if visited["0000"] {
		return -1
	}
	var queue []string
	queue = append(queue, "0000")
	visited["0000"] = true
	depth := -1
	for len(queue) > 0 {
		size := len(queue)
		depth++
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return depth
			}
			chars := []byte(cur)
			for j := 0; j < 4; j++ {
				for d := -1; d <= 1; d += 2 {
					tmp := chars[j]
					chars[j] = byte('0' + (int(chars[j]-'0')+d+10)%10)
					next := string(chars)
					if !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
					chars[j] = tmp
				}
			}
		}
	}
	return -1
}
