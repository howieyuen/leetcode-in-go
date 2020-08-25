package _332_reconstruct_itinerary

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {
	var m = make(map[string][]string)
	for _, ticket := range tickets {
		m[ticket[0]] = append(m[ticket[0]], ticket[1])
	}
	for key := range m {
		sort.Strings(m[key])
	}
	
	var ans []string
	var dfs func(start string)
	dfs = func(start string) {
		for {
			if ends, ok := m[start]; !ok || len(ends) == 0 {
				break
			}
			tmp := m[start][0]
			m[start] = m[start][1:]
			dfs(tmp)
		}
		ans = append(ans, start)
	}
	dfs("JFK")
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
