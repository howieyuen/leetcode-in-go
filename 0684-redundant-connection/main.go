package _684_redundant_connection

func findRedundantConnection(edges [][]int) []int {
	n := len(edges) + 1
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	for _, edge := range edges {
		a, b := find(parents, edge[0]), find(parents, edge[1])
		if a == b {
			return edge
		}
		parents[b] = a
	}
	return nil
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func findRedundantConnection1(edges [][]int) []int {
	u := newUnionFindSet(len(edges) + 1)
	for _, edge := range edges {
		if !u.union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
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
