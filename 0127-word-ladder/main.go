package _127_word_ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	step := 0
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	startUsed := make([]bool, len(wordList))
	endUsed := make([]bool, len(wordList))
	if i := indexFunc(endWord, wordList); i == -1 {
		return 0
	} else {
		endUsed[i] = true
	}
	for len(startQueue) > 0 {
		step++
		l := len(startQueue)
		for i := 0; i < l; i++ {
			for j, w := range wordList {
				if !startUsed[j] && compare(startQueue[i], w) {
					startQueue = append(startQueue, w)
					startUsed[j] = true
					if endUsed[j] {
						return step + 1
					}
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
			startUsed, endUsed = endUsed, startUsed
		}
	}
	return 0
}

func indexFunc(str string, bank []string) int {
	for i, s := range bank {
		if str == s {
			return i
		}
	}
	return -1
}

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	if !exists(endWord, wordList) {
		return 0
	}
	step := 0
	visited := make([]bool, len(wordList))
	queue := []string{beginWord}
	for len(queue) > 0 {
		step++
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == endWord {
				return step
			}
			for j, w := range wordList {
				if !visited[j] && compare(queue[i], w) {
					queue = append(queue, w)
					visited[j] = true
				}
			}
		}
		queue = queue[l:]
	}
	return 0
}

func exists(end string, wordList []string) bool {
	for i := range wordList {
		if wordList[i] == end {
			return true
		}
	}
	return false
}

func compare(start, end string) bool {
	diff := 0
	for i := 0; i < len(start) && i < len(end); i++ {
		if start[i] != end[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	return diff == 1
}
