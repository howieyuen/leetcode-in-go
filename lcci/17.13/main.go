package _7_13

func respace(dictionary []string, sentence string) int {
	var dp = make([]int, len(sentence)+1)
	for i := 1; i <= len(sentence); i++ {
		dp[i] = dp[i-1] + 1
		for _, word := range dictionary {
			if len(word) <= i {
				if sentence[i-len(word):i] == word {
					dp[i] = min(dp[i], dp[i-len(word)])
				}
			}
		}
	}
	return dp[len(sentence)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 字典树算法，减少重复查询dictionary
func respace1(dictionary []string, sentence string) int {
	var root = new(Trie)
	for _, word := range dictionary {
		root.insert(word)
	}
	n := len(sentence)
	var dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		indices := root.search(sentence, i-1)
		for _, idx := range indices {
			dp[i] = min(dp[i], dp[idx])
		}
	}
	return dp[n]
}

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

// 找到以sentence[index]结尾的单词，其首字母在sentence中的下标
func (t *Trie) search(sentence string, index int) []int {
	var indices []int
	cur := t
	for i := index; i >= 0; i-- {
		c := sentence[i] - 'a'
		if cur.next[c] == nil {
			break
		}
		cur = cur.next[c]
		if cur.isEnd {
			indices = append(indices, i)
		}
	}
	return indices
}

// 单词倒序插入字典树
func (t *Trie) insert(word string) {
	cur := t
	for i := len(word) - 1; i >= 0; i-- {
		c := word[i] - 'a'
		if cur.next[c] == nil {
			cur.next[c] = new(Trie)
		}
		cur = cur.next[c]
	}
	cur.isEnd = true
}
