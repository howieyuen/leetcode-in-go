package _047_remove_all_adjacent_duplicates_in_string

func removeDuplicates(S string) string {
	stack := make([]byte, 0)
	stack = append(stack, S[0])
	for i := 1; i < len(S); i++ {
		stack = append(stack, S[i])
		stackSize := len(stack)
		if stackSize > 1 && stack[stackSize-1] == stack[stackSize-2] {
			stack = stack[:stackSize-2]
		}
	}
	return string(stack)
}
