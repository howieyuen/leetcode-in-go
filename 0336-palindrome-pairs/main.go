package _336_palindrome_pairs

func palindromePairs(words []string) [][]int {
	check := make(map[string]int, len(words))
	for i := range words {
		check[words[i]] = i
	}
	
	var res [][]int
	for i, w := range words {
		m := len(w)
		for j := 0; j < m; j++ {
			if idx, ok := check[reverse(w[:j])]; ok && isPalindrome(w, j, m-1) && idx != i {
				res = append(res, []int{i, idx})
				if len(words[idx]) == 0 {
					res = append(res, []int{idx, i})
				}
			}
			if idx, ok := check[reverse(w[j:])]; ok && isPalindrome(w, 0, j-1) && idx != i {
				res = append(res, []int{idx, i})
				if len(words[idx]) == 0 {
					res = append(res, []int{i, idx})
				}
			}
		}
	}
	return res
}

func reverse(s string) string {
	res := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		res[i], res[len(s)-1-i] = res[len(s)-1-i], res[i]
	}
	return string(res)
}

var tree []Node

func palindromePairs1(words []string) [][]int {
	tree = []Node{{chars: [26]int{}, index: -1}}
	for i := range words {
		insert(words[i], i)
	}
	
	var ans [][]int
	for i := range words {
		m := len(words[i])
		for j := 0; j <= m; j++ {
			if isPalindrome(words[i], j, m-1) {
				left := findWord(words[i], 0, j-1)
				if left != -1 && left != i {
					ans = append(ans, []int{i, left})
				}
			}
			
			if j > 0 && isPalindrome(words[i], 0, j-1) {
				right := findWord(words[i], j, m-1)
				if right != -1 && right != i {
					ans = append(ans, []int{right, i})
				}
			}
		}
	}
	return ans
}

type Node struct {
	chars [26]int
	index int
}

func insert(word string, id int) {
	var add = 0
	for i := range word {
		x := word[i] - 'a'
		if tree[add].chars[x] == 0 {
			tree = append(tree, Node{chars: [26]int{}, index: -1})
			tree[add].chars[x] = len(tree) - 1
		}
		add = tree[add].chars[x]
	}
	tree[add].index = id
}

func findWord(word string, left, right int) int {
	var add = 0
	for i := right; i >= left; i-- {
		x := word[i] - 'a'
		if tree[add].chars[x] == 0 {
			return -1
		}
		add = tree[add].chars[x]
	}
	return tree[add].index
}

func isPalindrome(word string, left, right int) bool {
	for i := 0; i < (right-left+1)/2; i++ {
		if word[left+i] != word[right-i] {
			return false
		}
	}
	return true
}
