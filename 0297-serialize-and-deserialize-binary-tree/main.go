package _297_serialize_and_deserialize_binary_tree

import (
	"container/list"
	"fmt"
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
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return serialize(root)
}

func serialize(root *TreeNode) string {
	if root == nil {
		return "null,"
	}
	str := strconv.Itoa(root.Val) + ","
	str += serialize(root.Left)
	str += serialize(root.Right)
	return str
}

// Deserialize your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	values := strings.Split(data, ",")
	queue := list.New()
	for i := range values {
		queue.PushBack(values[i])
	}
	return buildTree(queue)
}

func buildTree(queue *list.List) *TreeNode {
	if queue.Len() == 0 {
		return nil
	}
	str := queue.Remove(queue.Front()).(string)
	if str == "null" {
		return nil
	}
	val, _ := strconv.Atoi(str)
	root := &TreeNode{Val: val}
	root.Left = buildTree(queue)
	root.Right = buildTree(queue)
	return root
}

func (c *Codec) serialize1(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := []*TreeNode{root}
	data := []string{}
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		if root != nil {
			data = append(data, strconv.Itoa(root.Val))
			queue = append(queue, root.Left, root.Right)
		} else {
			data = append(data, "null")
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(trimRightNulls(data, "null"), ","))
}

func trimRightNulls(data []string, target string) []string {
	right := len(data) - 1
	for data[right] == target {
		right--
	}
	return data[:right+1]
}

func (c *Codec) deserialize1(data string) *TreeNode {
	if data == "" {
		return nil
	}
	vals := strings.Split(data[1:len(data)-1], ",")
	if len(vals) == 0 {
		return nil
	}
	rootVal, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	i := 0
	for {
		if i++; i >= len(vals) {
			break
		}
		node := queue[0]
		queue = queue[1:]
		if vals[i] != "null" {
			leftVal, _ := strconv.Atoi(vals[i])
			node.Left = &TreeNode{Val: leftVal}
			queue = append(queue, node.Left)
		}
		
		if i++; i >= len(vals) {
			break
		}
		if vals[i] != "null" {
			rightVal, _ := strconv.Atoi(vals[i])
			node.Right = &TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
	}
	return root
}
