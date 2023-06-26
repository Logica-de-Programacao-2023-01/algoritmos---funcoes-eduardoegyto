func ordem(x []int) ([]int, error) {
	if len(x) == 0 {
		return nil, errors.New("O slice está vazio")
	}
	sorteio := make([]int, len(x))
	copy(sorteio, x)
	sort.Ints(sorteio)
	return sorteio, nil
}
