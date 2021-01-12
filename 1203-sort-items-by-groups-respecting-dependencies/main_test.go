package _203_sort_items_by_groups_respecting_dependencies

import (
	"reflect"
	"testing"
)

func Test_sortItems(t *testing.T) {
	type args struct {
		n           int
		m           int
		group       []int
		beforeItems [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n:           8,
				m:           2,
				group:       []int{-1, -1, 1, 0, 0, 1, 0, -1},
				beforeItems: [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}},
			},
			want: []int{6, 3, 4, 5, 2, 0, 7, 1},
		},
		{
			args: args{
				n:           8,
				m:           2,
				group:       []int{-1, -1, 1, 0, 0, 1, 0, -1},
				beforeItems: [][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortItems(tt.args.n, tt.args.m, tt.args.group, tt.args.beforeItems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func topSort(graph [][]int, deg, items []int) (orders []int) {
	q := []int{}
	for _, i := range items {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		orders = append(orders, from)
		for _, to := range graph[from] {
			deg[to]--
			if deg[to] == 0 {
				q = append(q, to)
			}
		}
	}
	return
}

func sortItems1(n, m int, group []int, beforeItems [][]int) (ans []int) {
	groupItems := make([][]int, m+n) // groupItems[i] 表示第 i 个组负责的项目列表
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i // 给不属于任何组的项目分配一个组
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	groupGraph := make([][]int, m+n) // 组间依赖关系
	groupDegree := make([]int, m+n)
	itemGraph := make([][]int, n) // 组内依赖关系
	itemDegree := make([]int, n)
	for cur, items := range beforeItems {
		curGroupID := group[cur]
		for _, pre := range items {
			preGroupID := group[pre]
			if preGroupID != curGroupID { // 不同组项目，确定组间依赖关系
				groupGraph[preGroupID] = append(groupGraph[preGroupID], curGroupID)
				groupDegree[curGroupID]++
			} else { // 同组项目，确定组内依赖关系
				itemGraph[pre] = append(itemGraph[pre], cur)
				itemDegree[cur]++
			}
		}
	}

	// 组间拓扑序
	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}
	groupOrders := topSort(groupGraph, groupDegree, items)
	if len(groupOrders) < len(items) {
		return nil
	}

	// 按照组间的拓扑序，依次求得各个组的组内拓扑序，构成答案
	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := topSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}
	return
}
