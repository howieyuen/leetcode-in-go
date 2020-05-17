package _210_course_schedule_ii

/*
图类型
根据节点出度，判断节点为叶子节点，自底向上
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 节点出度
	outDegree := make([]int, numCourses)
	// 边
	preCourses := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		outDegree[a]++
		if preCourses[b] == nil {
			preCourses[b] = []int{}
		}
		preCourses[b] = append(preCourses[b], a)
	}
	// 筛选当前叶子节点
	soloCourses := []int{}
	for i := 0; i < numCourses; i++ {
		if outDegree[i] == 0 {
			soloCourses = append(soloCourses, i)
		}
	}
	// 排序结果
	result := []int{}
	for len(soloCourses) != 0 {
		result = append(result, soloCourses[0])
		for _, course := range preCourses[soloCourses[0]] {
			outDegree[course]--
			// 节点入度为0，成为叶子
			if outDegree[course] == 0 {
				soloCourses = append(soloCourses, course)
			}
		}
		soloCourses = soloCourses[1:]
		numCourses--
	}
	if numCourses != 0 { // 未能遍历整个图
		return []int{}
	} else {
		return result
	}
}
