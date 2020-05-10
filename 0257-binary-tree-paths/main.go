package _257_binary_tree_paths

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	dfs(root, nil, &res)
	return res
}

func dfs(root *TreeNode, path []string, res *[]string) {
	path = append(path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(path, "->"))
	}
	if root.Left != nil {
		dfs(root.Left, path, res)
	}
	if root.Right != nil {
		dfs(root.Right, path, res)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
