package _809_expressive_words

func expressiveWords(S string, words []string) int {
	var ans int
	source := getExpress(S)
search:
	for _, word := range words {
		target := getExpress(word)
		if source.keys != target.keys {
			continue
		}

		for i := 0; i < len(target.counts); i++ {
			c1 := source.counts[i]
			c2 := target.counts[i]
			if c1 < 3 && c1 != c2 || c1 < c2 {
				continue search
			}
		}
		ans++
	}
	return ans
}

type Express struct {
	keys   string
	counts []int
}

func getExpress(s string) *Express {
	var keys []byte
	var counts []int
	prev := -1
	for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			keys = append(keys, s[i])
			counts = append(counts, i-prev)
			prev = i
		}
	}
	return &Express{
		keys:   string(keys),
		counts: counts,
	}
}
