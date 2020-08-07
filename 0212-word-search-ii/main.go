package _212_word_search_ii

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return nil
	}
	
	trie := new(Trie)
	for _, word := range words {
		trie.insert(word)
	}
	
	var ans []string
	
	rows, cols := len(board), len(board[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	var dfs func(row, col int, trie *Trie)
	dfs = func(row, col int, trie *Trie) {
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return
		}
		
		char := board[row][col]
		if char == ' ' || trie.next[char-'a'] == nil {
			return
		}
		
		trie = trie.next[char-'a']
		if trie.word != "" {
			ans = append(ans, trie.word)
			trie.word = "" // 避免重复添加
		}
		
		board[row][col] = ' '
		for _, dir := range directions {
			dfs(row+dir[0], col+dir[1], trie)
		}
		board[row][col] = char
	}
	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, trie)
		}
	}
	return ans
}

type Trie struct {
	next [26]*Trie
	word string
}

func (t *Trie) insert(word string) {
	for i := range word {
		index := word[i] - 'a'
		if t.next[index] == nil {
			t.next[index] = new(Trie)
		}
		t = t.next[index]
	}
	t.word = word
}
