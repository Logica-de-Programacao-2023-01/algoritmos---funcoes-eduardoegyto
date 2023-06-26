func aplicacao(x []int, fn func(int) int) ([]int, error) {
	if len(x) == 0 {
		return nil, errors.New("O slice está vazio")
	}
	resultado := make([]int, len(x))
	for i, num := range x {
		resultado[i] = fn(num)
	}
	return resultado, nil
}
