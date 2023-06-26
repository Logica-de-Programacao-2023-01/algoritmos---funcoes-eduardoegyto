func par(x []int) ([]int, error) {
	if len(x) == 0 {
		return nil, errors.New("O slice está vazio")
	}
	var pares []int
	for _, num := range x {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}
	return pares, nil
}
