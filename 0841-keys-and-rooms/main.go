package _841_keys_and_rooms

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	visited := make([]bool, len(rooms))
	visited[0] = true
	currentKeys := rooms[0]
	for len(currentKeys) != 0 {
		var newKeys []int
		for _, k := range currentKeys {
			if !visited[k] {
				visited[k] = true
				newKeys = append(newKeys, rooms[k]...)
			}
		}
		currentKeys = newKeys
	}
	for _, v := range visited {
		if v == false {
			return false
		}
	}
	return true
}
