package _990_satisfiability_of_equality_equations


func equationsPossible(equations []string) bool {
	var values = make([]int, 26)
	for i := 0; i < 26; i++ {
		values[i] = i
	}
	for i := range equations {
		if equations[i][1] == '=' {
			index1 := int(equations[i][0] - 'a')
			index2 := int(equations[i][3] - 'a')
			union(values, index1, index2)
		}
	}
	for i := range equations {
		if equations[i][1] == '!' {
			index1 := int(equations[i][0] - 'a')
			index2 := int(equations[i][3] - 'a')
			if find(values, index1) == find(values, index2) {
				return false
			}
		}
	}
	return true
}

func union(values []int, index1, index2 int) {
	values[find(values, index1)] = find(values, index2)
}

func find(values []int, index int) int {
	for values[index] != index {
		values[index] = values[values[index]]
		index = values[index]
	}
	return index
}
