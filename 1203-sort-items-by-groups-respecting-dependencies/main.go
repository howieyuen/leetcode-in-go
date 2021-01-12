package _203_sort_items_by_groups_respecting_dependencies

// n: 项目数量
// m: 小组数量
// group: group[i]=j 表示项目i由小组j接手
// beforeItems: beforeItems[i]表示必须在项目i之间的项目集合
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// groupItems[i]表示i组负责的项目集合
	groupItems := make([][]int, m+n)
	for i := range group {
		if group[i] == -1 {
			// 数据预处理，当前无人接手的项目分配一个假组
			group[i] = m + i
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	// 组之间依赖关系
	groupGraph := make([][]int, m+n)
	// 组的优先级
	groupDegree := make([]int, m+n)
	// 组内项目的依赖关系
	itemGraph := make([][]int, n)
	// 项目优先级
	itemDegree := make([]int, n)
	for cur, items := range beforeItems {
		// 当前项目所属的组
		curGroup := group[cur]
		for _, pre := range items {
			// 必须在当前项目之前的项目所属的组
			preGroup := group[pre]
			if preGroup != curGroup { // 不同项目组
				// 组之间关系
				groupGraph[preGroup] = append(groupGraph[preGroup], curGroup)
				// 当前项目所属组入度+1
				groupDegree[curGroup]++
			} else { // 相同组
				// 项目之间关系
				itemGraph[pre] = append(itemGraph[pre], cur)
				// 当前项目入度+1
				itemDegree[cur]++
			}
		}
	}

	// 组 拓扑排序
	groups := make([]int, m+n)
	for i := range groups {
		groups[i] = i
	}
	groupOrders := topoSort(groupGraph, groupDegree, groups)
	if len(groupOrders) < len(groups) {
		return nil
	}

	var ans []int
	// 组内项目拓扑排序
	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := topoSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}
	return ans
}

// 拓扑排序
func topoSort(graph [][]int, degree, items []int) []int {
	var queue []int
	for _, item := range items {
		// 入度为0的节点优先
		if degree[item] == 0 {
			queue = append(queue, item)
		}
	}
	var orders []int
	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]
		orders = append(orders, from)
		for _, to := range graph[from] {
			// 摘除from节点，to节点入度-1
			degree[to]--
			if degree[to] == 0 {
				queue = append(queue, to)
			}
		}
	}
	return orders
}
