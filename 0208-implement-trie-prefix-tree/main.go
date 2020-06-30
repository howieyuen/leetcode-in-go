package _208_implement_trie_prefix_tree

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for i := range word {
		index := word[i] - 'a'
		if this.children[index] == nil {
			this.children[index] = new(Trie)
		}
		this = this.children[index]
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, b := range word {
		if this.children[b-'a'] == nil {
			return false
		}
		this = this.children[b-'a']
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, b := range prefix {
		if this.children[b-'a'] == nil {
			return false
		}
		this = this.children[b-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
