package test2

func getWinner(arr []int, k int) int {
	i, t := 0, 0
	for ; i < len(arr)-1 && t < k; i++ {
		if arr[i] > arr[i+1] {
			arr[i+1] = arr[i]
			t++
		} else {
			t = 1
		}
	}
	return arr[i]
}
