package _464_can_i_win

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	// 第一种特殊情况，第一个人选一次即可到达预期值
	if desiredTotal <= 0 {
		return true
	}
	// 第二种特殊情况，所有元素的和都小于预期值，则永远无法赢
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	memo := make([]int, 1<<maxChoosableInteger)
	return canWin(maxChoosableInteger, desiredTotal, memo, 0)
}

// max<=20，所以可以使用一个int型的各个位标记是否使用
// memo[state]用于标记在使用state的情况本次挑选是否能赢（二进制各个位真值代表某个元素是否已经使用，state = 1101代表使用了1，3，4）
func canWin(max int, desired int, memo []int, state int) bool {
	if desired <= 0 {
		return false
	}
	if memo[state] == 1 {
		return true
	}
	for i := 0; i < max; i++ {
		// 第i位表示选择[1,2,3, max]选择i + 1这个值
		cur := 1 << i
		// 使用过
		if state&cur != 0 {
			continue
		}
		// desired - (i + 1)表示选择了i + 1，期待值减小
		// cur|state代表更新各个元素使用情况，使用i + 1，将used的第i位（从低到高）标记为1
		// !canWin()表示的是对方选择输了
		if !canWin(max, desired-i-1, memo, state|cur) {
			memo[state] = 1
			return true
		}
	}
	memo[state] = -1
	return false
}
