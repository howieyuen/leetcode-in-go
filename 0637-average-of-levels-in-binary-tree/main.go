package _637_average_of_levels_in_binary_tree

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var ans []float64
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		sum := 0.0
		for i := 0; i < size; i++ {
			sum += float64(queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		ans = append(ans, sum/float64(size))
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
