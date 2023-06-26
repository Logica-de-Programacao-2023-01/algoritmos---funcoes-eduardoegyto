func posicao(x []int, valor int) int {
	for i, num := range x {
		if num == valor {
			return i
		}
	}
	return -1
}
