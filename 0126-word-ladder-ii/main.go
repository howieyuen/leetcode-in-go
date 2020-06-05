package _126_word_ladder_ii

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordMap := map[string]struct{}{}
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	
	if _, ok := wordMap[endWord]; !ok {
		return [][]string{}
	}
	
	src := map[string]struct{}{beginWord: {}}
	dst := map[string]struct{}{endWord: {}}
	backward := false
	paths := map[string][]string{}
	
	for len(src) > 0 && len(dst) > 0 {
		if len(src) > len(dst) {
			src, dst = dst, src
			backward = !backward
		}
		for word := range src {
			delete(wordMap, word)
		}
		found := false
		newSrc := map[string]struct{}{}
		for word := range src {
			wordBytes := []byte(word)
			for i := range wordBytes {
				for c := byte('a'); c <= 'z'; c++ {
					if word[i] == c {
						continue
					}
					wordBytes[i] = c
					target := string(wordBytes)
					if _, ok := wordMap[target]; !ok {
						continue
					}
					source := word
					if _, ok := dst[target]; ok {
						found = true
					} else {
						newSrc[target] = struct{}{}
					}
					if backward {
						source, target = target, source
					}
					paths[target] = append(paths[target], source)
				}
				wordBytes[i] = word[i]
			}
		}
		if found {
			break
		}
		src = newSrc
	}
	
	ans := [][]string{}
	var getPath func(cur string, path []string)
	getPath = func(cur string, path []string) {
		if cur == beginWord {
			newPath := make([]string, len(path))
			copy(newPath, path)
			ans = append(ans, newPath)
			return
		}
		srcs, ok := paths[cur]
		if !ok {
			return
		}
		path = append([]string{""}, path...)
		for _, src := range srcs {
			path[0] = src
			getPath(src, path)
		}
	}
	getPath(endWord, []string{endWord})
	return ans
}
