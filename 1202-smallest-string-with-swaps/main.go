package _202_smallest_string_with_swaps

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	u := newUnionSet(n)
	res := make([]byte, n)
	for _, p := range pairs {
		u.union(p[0], p[1])
	}

	// 格式化并查集，哪些下标可以交换
	set := make(map[int][]int)
	for i := 0; i < n; i++ {
		set[u.find(i)] = append(set[u.find(i)], i)
	}

	// 同一并查集按字典序排序
	for _, originIndexes := range set {
		// 记录交换前的下标
		targetIndexes := make([]int, len(originIndexes))
		copy(targetIndexes, originIndexes)
		// 排序
		sort.Slice(originIndexes, func(i, j int) bool {
			return s[originIndexes[i]] < s[originIndexes[j]]
		})
		// 赋值
		for i := 0; i < len(originIndexes); i++ {
			res[targetIndexes[i]] = s[originIndexes[i]]
		}
	}

	return string(res)
}

type unionSet struct {
	// 记录节点的父节点，相同父节点属于同一个连通图
	parents []int
	// 节点的秩，主要记录节点位于树的深度，从叶子节点出发
	// 秩越小，越“儿子”
	// 秩越大，越“父亲”
	ranks   []int
}

func newUnionSet(n int) *unionSet {
	parents := make([]int, n)
	ranks := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	return &unionSet{parents: parents, ranks: ranks}
}

// 路径压缩
func (u *unionSet) find(x int) int {
	if u.parents[x] != x {
		u.parents[x] = u.find(u.parents[x])
	}
	return u.parents[x]
}

// 合并2个节点
// 同一个并查集，判断px和py的秩大小，
func (u *unionSet) union(x, y int) {
	px, py := u.find(x), u.find(y)
	if px != py {
		if u.ranks[px] < u.ranks[py] {
			px, py = py, px
		}
		// py的父节点是px
		u.parents[py] = px
		// 更新秩
		if u.ranks[px] == u.ranks[py] {
			u.ranks[px] += 1
		}
	}
}
