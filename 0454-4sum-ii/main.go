package _454_4sum_ii

func fourSumCount(A []int, B []int, C []int, D []int) int {
	result := 0
	sum := make(map[int]int)
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			sum[C[i]+D[j]]++
		}
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			result += sum[-A[i]-B[j]]
			
		}
	}
	return result
}

func fourSumCount1(A []int, B []int, C []int, D []int) int {
	maps1 := generateMap(A, B)
	maps2 := generateMap(C, D)
	ans := 0
	for key, val1 := range maps1 {
		if val2, ok := maps2[-key]; ok {
			ans += val1 * val2
		}
	}
	return ans
	/*ans := 0
	input := map[int][]int{
		0: A,
		1: B,
		2: C,
		3: D,
	}
	visited := map[int]bool{
		0: false,
		1: false,
		2: false,
		3: false,
	}
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			a := input[i]
			visited[i] = true
			b := input[j]
			visited[j] = true
			var c, d []int
			for k, v := range input {
				if !visited[k] {
					c = v
					d = input[6-i-j-k]
					break
				}
			}
			ans += twoSumCount(a, b, c, d)
			visited[i] = false
			visited[j] = false
		}
	}
	return ans*/
}

func generateMap(A []int, B []int) map[int]int {
	var maps = make(map[int]int)
	for i := range A {
		for j := range B {
			key := A[i] + B[j]
			if _, ok := maps[key]; ok {
				maps[key]++
			} else {
				maps[key] = 1
			}
		}
	}
	return maps
}
