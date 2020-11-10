package _449_serialize_and_deserialize_bst

import (
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var values []string
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root != nil {
			values = append(values, strconv.Itoa(root.Val))
			preOrder(root.Left)
			preOrder(root.Right)
		}
	}
	preOrder(root)
	return strings.Join(values, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	preOrderStr := strings.Split(data, ",")
	var preOrder = make([]int, len(preOrderStr))
	for i := range preOrderStr {
		preOrder[i], _ = strconv.Atoi(preOrderStr[i])
	}
	inOrder := make([]int, len(preOrder))
	copy(inOrder, preOrder)
	sort.Ints(inOrder)
	
	var buildTree func(preOrder, inOrder []int) *TreeNode
	buildTree = func(preOrder, inOrder []int) *TreeNode {
		if len(preOrder) == 0 {
			return nil
		}
		rootVal := preOrder[0]
		index := 0
		for index < len(inOrder) && inOrder[index] != rootVal {
			index++
		}
		root := &TreeNode{Val: rootVal}
		root.Left = buildTree(preOrder[1:index+1], inOrder[:index])
		root.Right = buildTree(preOrder[index+1:], inOrder[index+1:])
		return root
	}
	return buildTree(preOrder, inOrder)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
