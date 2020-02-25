package _079_word_search

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.
*/
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, word string, k int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	tmp := board[i][j]
	board[i][j] = ' '
	flag := dfs(board, i+1, j, word, k+1) ||
		dfs(board, i-1, j, word, k+1) ||
		dfs(board, i, j+1, word, k+1) ||
		dfs(board, i, j-1, word, k+1)
	board[i][j] = tmp
	return flag
}
