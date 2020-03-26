package main

import (
	`fmt`
)

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

func main() {
	var s = []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, x := range s {
		fmt.Println(isValid(x))
	}
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var stack = NewStack()
	for _, ch := range s {
		switch ch {
		case '}':
			if stack.isEmpty() || stack.pop() != '{' {
				return false
			} else {
				continue
			}
		case ']':
			if stack.isEmpty() || stack.pop() != '[' {
				return false
			} else {
				continue
			}
		case ')':
			if stack.isEmpty() || stack.pop() != '(' {
				return false
			} else {
				continue
			}
		default:
			stack.push(ch)
		}
	}
	return stack.isEmpty()
}

type Stack struct {
	array []interface{}
}

func NewStack() *Stack {
	return &Stack{
		array: make([]interface{}, 0),
	}
}

func (queue *Stack) push(val interface{}) {
	queue.array = append(queue.array, val)
}

func (queue *Stack) pop() interface{} {
	val := queue.array[len(queue.array)-1]
	queue.array = queue.array[:len(queue.array)-1]
	return val
}

func (queue *Stack) isEmpty() bool {
	return len(queue.array) == 0
}
