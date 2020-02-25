package _022_Generate_Parentheses

func generateParenthesis(n int) []string {
	var ans []string
	dfs("", n, n, &ans)
	return ans
}

func dfs(cur string, left, right int, ans *[]string) {
	if left == 0 && right == 0 {
		*ans = append(*ans, cur)
		return
	}
	if left > 0 {
		dfs(cur+"(", left-1, right, ans)
	}
	if left < right {
		dfs(cur+")", left, right-1, ans)
	}
}