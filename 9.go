func palavras(x string) ([]string, error) {
	if x == "" {
		return nil, errors.New("A string está vazia")
	}
	palavra := strings.Fields(x)
	return words, nil
}

