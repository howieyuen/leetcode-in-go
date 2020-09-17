package _685_redundant_connection_ii

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges) + 1
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	u := newUnionFindSet(n)
	var conflictEdge, cycleEdge []int
	for _, e := range edges {
		from, to := e[0], e[1]
		if parent[to] != to {
			conflictEdge = e
		} else {
			parent[to] = from
			if u.find(from) == u.find(to) {
				cycleEdge = e
			} else {
				u.union(from, to)
			}
		}
	}
	if conflictEdge == nil {
		return cycleEdge
	}
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

type unionFindSet struct {
	ancestor []int
}

func newUnionFindSet(n int) *unionFindSet {
	ancestor := make([]int, n)
	for i := range ancestor {
		ancestor[i] = i
	}
	return &unionFindSet{ancestor: ancestor}
}

func (u *unionFindSet) find(x int) int {
	if u.ancestor[x] != x {
		u.ancestor[x] = u.find(u.ancestor[x])
	}
	return u.ancestor[x]
}

func (u *unionFindSet) union(from, to int) bool {
	from, to = u.find(from), u.find(to)
	if from == to {
		return false
	}
	u.ancestor[to] = from
	return true
}

func findRedundantDirectedConnection1(edges [][]int) []int {
	n := len(edges) + 1
	inDegree := make([]int, n)
	for _, e := range edges {
		inDegree[e[1]]++
	}
	
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 2 {
			if !judgeCycle(edges, n, i) {
				return edges[i]
			}
		}
	}
	
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 1 {
			if !judgeCycle(edges, n, i) {
				return edges[i]
			}
		}
	}
	
	return nil
}

func judgeCycle(edges [][]int, n int, removedIndex int) bool {
	u := newUnionFindSet(n)
	for i, e := range edges {
		if i == removedIndex {
			continue
		}
		if !u.union(e[0], e[1]) {
			return true
		}
	}
	return false
}
