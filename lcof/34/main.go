package _4

func pathSum(root *TreeNode, sum int) [][]int {
	var paths [][]int
	if root == nil {
		return paths
	}
	var dfs func(root *TreeNode, curSum int, path []int)
	dfs = func(root *TreeNode, curSum int, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		curSum += root.Val
		if curSum == sum && root.Left == nil && root.Right == nil {
			var tmp = make([]int, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
			return
		}
		if root.Left != nil {
			dfs(root.Left, curSum, path)
		}
		if root.Right != nil {
			dfs(root.Right, curSum, path)
		}
		curSum -= root.Val
		path = path[:len(path)-1]
	}
	dfs(root, 0, nil)
	return paths
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
