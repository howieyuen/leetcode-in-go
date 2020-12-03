package _710_random_pick_with_blacklist

import (
	"math/rand"
)

type Solution struct {
	whiteList map[int]int
	base      int
}

// 随机数在[0,N)之间，blacklist的长度len小于N，
// 则共有N-len个数是白名单，len个数是黑名单
// 从[0,N-len)之间选择随机数 x，在黑名单的数字映射到白名单即可
func Constructor(N int, blacklist []int) Solution {
	base := N - len(blacklist)
	whitelist := make(map[int]int, len(blacklist))
	for _, b := range blacklist {
		whitelist[b] = b
	}
	last := N - 1
	for _, b := range blacklist {
		if b >= base {
			continue
		}
		_, ok := whitelist[last]
		for ok {
			last--
			_, ok = whitelist[last]
		}
		whitelist[b] = last
		last--
	}
	return Solution{
		whiteList: whitelist,
		base:      base,
	}
}

func (this *Solution) Pick() int {
	x := rand.Intn(this.base)
	if y, ok := this.whiteList[x]; ok {
		return y
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
