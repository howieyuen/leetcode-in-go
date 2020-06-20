package _113_path_sum_ii

func pathSum(root *TreeNode, sum int) [][]int {
	var ans [][]int
	var path []int
	dfs(root, sum, path, &ans)
	return ans
}

func dfs(root *TreeNode, target int, path []int, ans *[][]int) {
	if root == nil {
		return
	}
	target -= root.Val
	path = append(path, root.Val)
	if target == 0 && root.Left == nil && root.Right == nil {
		var tmp = make([]int, len(path))
		copy(tmp, path)
		*ans = append(*ans, tmp)
		return
	}
	if root.Left != nil {
		dfs(root.Left, target, path, ans)
	}
	if root.Right != nil {
		dfs(root.Right, target, path, ans)
	}
	target += root.Val
	path = path[:len(path)-1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
