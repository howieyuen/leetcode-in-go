package _210_course_schedule_ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	outDegree := make([]int, numCourses)
	preCourses := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		outDegree[a]++
		if preCourses[b] == nil {
			preCourses[b] = []int{}
		}
		preCourses[b] = append(preCourses[b], a)
	}
	soloCourses := []int{}
	for i := 0; i < numCourses; i++ {
		if outDegree[i] == 0 {
			soloCourses = append(soloCourses, i)
		}
	}
	result := []int{}
	for len(soloCourses) != 0 {
		result = append(result, soloCourses[0])
		for _, course := range preCourses[soloCourses[0]] {
			outDegree[course]--
			if outDegree[course] == 0 {
				soloCourses = append(soloCourses, course)
			}
		}
		soloCourses = soloCourses[1:]
		numCourses--
	}
	if numCourses != 0 {
		return []int{}
	} else {
		return result
	}
}
