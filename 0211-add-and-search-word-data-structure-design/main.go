package _211_add_and_search_word_data_structure_design

type WordDictionary struct {
	root *Node
}

type Node struct {
	next [26]*Node
	end  bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{
			next: [26]*Node{},
			end:  false,
		},
	}
}

/** Adds a word into the data structure. */
func (wd *WordDictionary) AddWord(word string) {
	n := len(word)
	cur := wd.root
	for i := 0; i < n; i++ {
		index := word[i] - 'a'
		if cur.next[index] == nil {
			cur.next[index] = &Node{
				next: [26]*Node{},
				end:  false,
			}
		}
		cur = cur.next[index]
	}
	cur.end = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (wd *WordDictionary) Search(word string) bool {
	var dfs func(root *Node, word string) bool
	dfs = func(cur *Node, word string) bool {
		if cur == nil {
			return false
		}
		n := len(word)
		for i := 0; i < n; i++ {
			if word[i] == '.' {
				for j := 0; j < 26; j++ {
					if dfs(cur.next[j], word[i+1:]) {
						return true
					}
				}
				return false
			}
			nextIndex := word[i] - 'a'
			if cur.next[nextIndex] == nil {
				return false
			}
			cur = cur.next[nextIndex]
		}
		return cur.end
	}
	return dfs(wd.root, word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
