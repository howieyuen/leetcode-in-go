package _337_house_robber_iii

// 当前节点选择不偷：左孩子能偷到最多的钱 + 右孩子能偷到最多的钱
// 当前节点选择偷：左孩子选择自己不偷时能得到的钱 + 右孩子选择不偷时能得到的钱 + 当前节点的钱数
// res = [2]int{}, res[0]表示不取，res[1]表示取
func rob2(root *TreeNode) int {
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {
		var res = make([]int, 2)
		if root == nil {
			return res
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		res[0] = max(left[0], left[1]) + max(right[0], right[1])
		res[1] = root.Val + left[0] + right[0]
		return res
	}
	res := dfs(root)
	return max(res[0], res[1])
}

// 记忆法保存中间结果，避免重复计算
func rob1(root *TreeNode) int {
	var dfs func(root *TreeNode, memo map[*TreeNode]int) int
	dfs = func(root *TreeNode, memo map[*TreeNode]int) int {
		if root == nil {
			return 0
		}
		if money, ok := memo[root]; ok {
			return money
		}
		var money = root.Val
		if root.Left != nil {
			money += dfs(root.Left.Left, memo) + dfs(root.Left.Right, memo)
		}
		if root.Right != nil {
			money += dfs(root.Right.Left, memo) + dfs(root.Right.Right, memo)
		}
		money = max(money, dfs(root.Left, memo)+dfs(root.Right, memo))
		memo[root] = money
		return money
	}
	return dfs(root, make(map[*TreeNode]int))
}

// 隔层取钱
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var money = root.Val
	if root.Left != nil {
		money += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		money += rob(root.Right.Left) + rob(root.Right.Right)
	}
	return max(money, rob(root.Left)+rob(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
