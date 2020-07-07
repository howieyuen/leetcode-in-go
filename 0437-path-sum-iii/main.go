package _437_path_sum_iii

func pathSum1(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return dfs(root, sum) + pathSum1(root.Left, sum) + pathSum1(root.Right, sum)
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	if root.Val == sum {
		cnt++
	}
	return cnt + dfs(root.Left, sum-root.Val) + dfs(root.Right, sum-root.Val)
}

func pathSum(root *TreeNode, sum int) int {
	prefixCount := make(map[int]int)
	prefixCount[0] = 1
	return recursive(root, prefixCount, 0, sum)
}

func recursive(root *TreeNode, prefixCount map[int]int, cur, target int) int {
	if root == nil {
		return 0
	}
	res := 0
	cur += root.Val
	res += prefixCount[cur-target]
	prefixCount[cur]++
	res += recursive(root.Left, prefixCount, cur, target)
	res += recursive(root.Right, prefixCount, cur, target)
	prefixCount[cur]--
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
