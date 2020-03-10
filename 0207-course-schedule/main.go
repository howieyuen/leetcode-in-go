package _207_course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	outDegree := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		outDegree[prerequisites[i][0]]++
	}
	soloCourses := []int{}
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
