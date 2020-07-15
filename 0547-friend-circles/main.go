package _547_friend_circles

// 并查集
func findCircleNum2(M [][]int) int {
	u := NewUnion(len(M))
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 {
				u.union(i, j)
			}
		}
	}
	return u.count
}

type Union struct {
	count  int   // 初始节点个数
	parent []int // 父节点
	size   []int // 独立子树的节点总数
}

func NewUnion(n int) *Union {
	union := &Union{
		count:  n,
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		union.parent[i] = i // 初始父节点指向自身
		union.size[i] = 1
	}
	return union
}

func (u *Union) union(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)
	if rootP == rootQ {
		return
	}
	// 将size较小的子树连接到size较大的子树上
	if u.size[rootP] < u.size[rootQ] {
		u.parent[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	} else {
		u.parent[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	}
	u.count--
}

func (u *Union) find(x int) int {
	for x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

// bfs
func findCircleNum1(M [][]int) int {
	var visited = make([]bool, len(M))
	var queue []int
	var ans int
	for i := 0; i < len(M); i++ {
		if !visited[i] {
			queue = append(queue, i)
			for len(queue) > 0 {
				s := queue[0]
				queue = queue[1:]
				visited[s] = true
				for j := 0; j < len(M); j++ {
					if M[s][j] == 1 && !visited[j] {
						queue = append(queue, j)
					}
				}
			}
			ans++
		}
	}
	return ans
}

// dfs
func findCircleNum(M [][]int) int {
	var visited = make([]bool, len(M))
	var dfs func(i int)
	dfs = func(i int) {
		for j := range visited {
			if M[i][j] == 1 && !visited[j] {
				visited[j] = true
				dfs(j)
			}
		}
	}
	var ans int
	for i := range M {
		if !visited[i] {
			dfs(i)
			ans++
		}
	}
	return ans
}
