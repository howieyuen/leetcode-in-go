package _236_lowest_common_ancestor_of_a_binary_tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	visited := make(map[int]bool)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parent[root.Left.Val] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			parent[root.Right.Val] = root
			dfs(root.Right)
		}
	}
	dfs(root)
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
