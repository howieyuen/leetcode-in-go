package _207_course_schedule

// 校验出度合法
func canFinish(numCourses int, prerequisites [][]int) bool {
	outDegree := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		outDegree[prerequisites[i][0]]++
	}
	var soloCourses []int
	for i := 0; i < numCourses; i++ {
		if outDegree[i] == 0 {
			soloCourses = append(soloCourses, i)
		}
	}
	for len(soloCourses) > 0 {
		for i := 0; i < len(prerequisites); i++ {
			if prerequisites[i][1] != soloCourses[0] {
				continue
			}
			outDegree[prerequisites[i][0]]--
			if outDegree[prerequisites[i][0]] == 0 {
				soloCourses = append(soloCourses, prerequisites[i][0])
			}
		}
		soloCourses = soloCourses[1:]
		numCourses--
	}
	return numCourses == 0
}

// dfs
func canFinish1(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
	)
	
	var dfs func(u int)
	var valid = true
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}
	
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// bfs
func canFinish2(numCourses int, prerequisites [][]int) bool {
	var (
		edges     = make([][]int, numCourses)
		outDegree = make([]int, numCourses)
		result    []int
	)
	
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		outDegree[info[0]]++
	}
	
	var queue []int
	for i := 0; i < numCourses; i++ {
		if outDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			outDegree[v]--
			if outDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return len(result) == numCourses
}