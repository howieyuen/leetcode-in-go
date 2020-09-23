package _968_binary_tree_cameras

// 状态1：放摄像头
// 状态2：不放摄像头，但是可以被监测到
// 状态3：监测不到
func minCameraCover(root *TreeNode) int {
	const (
		withCam = iota
		withInWatch
		beyondWatch
	)
	var count int
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return withInWatch
		}
		leftStatus := dfs(root.Left)
		rightStatus := dfs(root.Right)
		if leftStatus == beyondWatch || rightStatus == beyondWatch {
			count++
			return withCam
		}
		if leftStatus == withCam || rightStatus == withCam {
			return withInWatch
		}
		return beyondWatch
	}
	if dfs(root) == beyondWatch {
		count++
	}
	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
