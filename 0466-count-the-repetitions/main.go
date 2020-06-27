package _466_count_the_repetitions

// n1个s1组成S1，n2个s2组成S2；S1与s2匹配，计算出S1中包含s2的个数cnt，cnt/n2即为结果
// 暴力计算即为挨个匹配，要加速运算过程，只需要保存匹配过程：
// 记录每匹配完一个s1，保存s1和s2的下标对应关系，循环中重复计算过程可直接跳过
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	len1, len2 := len(s1), len(s2)
	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}
	// s1与s2移位匹配过程中，s1的每个字符与index1、s2的每个字符与index2的对应关系
	map1, map2 := make(map[int]int), make(map[int]int)
	cnt := 0
	for index1, index2 := 0, 0; index1 < len1*n1; index1++ {
		if index1%len1 == len1-1 { // 已到s1末位
			if lastIndex1, ok := map1[index2%len2]; ok { // 此时s2位置已出现，找到重复比较部分
				// 每个环需要s1的个数
				cycleS1Num := (index1 - lastIndex1) / len1
				// 剩余环数
				leftCycles := (n1 - 1 - index1/len1) / cycleS1Num
				// 每个环含有s2的个数
				cycleS2Num := index2/len2 - map2[index2%len2]/len2
				// 将S1快进到相应的位置
				index1 += leftCycles * cycleS1Num * len1
				// 把快进部分的答案数量加上
				cnt += leftCycles * cycleS2Num
			} else {
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}
		}
		// 字符匹配，index2++
		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				cnt++
			}
			index2++
		}
	}
	return cnt / n2
}

func getMaxRepetitions1(s1 string, n1 int, s2 string, n2 int) int {
	cnt := 0
	p2 := 0
	for i := 0; i < n1; i++ {
		for p1 := 0; p1 < len(s1); p1++ {
			if s1[p1] == s2[p2] {
				p2++
				if p2 == len(s2) {
					cnt++
					p2 = 0
				}
			}
		}
	}
	return cnt / n2
}
