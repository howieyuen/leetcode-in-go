package _096_brace_expansion_ii

import (
	"sort"
	"strings"
)

func braceExpansionII(expression string) []string {
	var groups [][][]string
	groups = append(groups, [][]string{})
	level := 0
	start := -1
	for i, c := range expression {
		if c == '{' {
			if level == 0 {
				start = i + 1
			}
			level++
		} else if c == '}' {
			level--
			if level == 0 {
				subAns := braceExpansionII(expression[start:i])
				groups[len(groups)-1] = append(groups[len(groups)-1], subAns)
			}
		} else if c == ',' && level == 0 {
			groups = append(groups, [][]string{})
		} else if level == 0 {
			groups[len(groups)-1] = append(groups[len(groups)-1], []string{string(c)})
		}
	}
	
	set := map[string]bool{}
	var ans []string
	for _, group := range groups {
		pre := []string{""}
		for _, g := range group {
			var tmp []string
			for _, a := range pre {
				for _, b := range g {
					tmp = append(tmp, a+b)
				}
			}
			pre = tmp
		}
		for _, p := range pre {
			if !set[p] {
				set[p] = true
				ans = append(ans, p)
			}
		}
	}
	sort.Strings(ans)
	
	return ans
}

func braceExpansionII1(expression string) []string {
	var queue []string
	queue = append(queue, expression)
	
	var ans []string
	var set = map[string]bool{}
	for len(queue) > 0 {
		exp := queue[0]
		queue = queue[1:]
		if strings.IndexByte(exp, '{') == -1 {
			if !set[exp] {
				set[exp] = true
				ans = append(ans, exp)
			}
			continue
		}
		i := 0
		var lStart, rStart int
		for exp[i] != '}' {
			if exp[i] == '{' {
				lStart = i
			}
			i++
		}
		rStart = i
		lStr, rStr := exp[:lStart], exp[rStart+1:]
		subs := strings.Split(exp[lStart+1:rStart], ",")
		for _, s := range subs {
			queue = append(queue, lStr+s+rStr)
		}
	}
	sort.Strings(ans)
	return ans
}

func braceExpansionII2(s string) []string {
	res := dfs(s)
	sort.Strings(res)
	i, j := 0, 1
	for j < len(res) {
		if res[j] != res[i] {
			i++
			res[i] = res[j]
		}
		j++
	}
	return res[:i+1]
}

func dfs(s string) []string {
	if !strings.Contains(s, "{") {
		return []string{s}
	}
	
	start := 0
	for s[start] != '{' {
		start++
	}
	if start > 0 {
		return merge([]string{s[:start]}, dfs(s[start:]))
	}
	
	end := start
	depth := 1
	for depth != 0 {
		end++
		if s[end] == '{' {
			depth++
		}
		if s[end] == '}' {
			depth--
		}
	}
	
	var list []string
	for l, r := start+1, start+1; r <= end; r++ {
		if s[r] == '{' {
			depth++
		}
		if s[r] == '}' {
			depth--
		}
		if (s[r] == ',' && depth == 0) || r == end {
			list = append(list, dfs(s[l:r])...)
			l = r + 1
		}
	}
	return merge(list, dfs(s[end+1:]))
}

func merge(list1, list2 []string) []string {
	var res []string
	for _, s1 := range list1 {
		for _, s2 := range list2 {
			res = append(res, s1+s2)
		}
	}
	return res
}
