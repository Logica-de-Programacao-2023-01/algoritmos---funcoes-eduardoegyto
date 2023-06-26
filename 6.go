func concatenacao(x []string) (string, error) {
	if len(x) == 0 {
		return "", errors.New("O slice está vazio")
	}
	return x.Join(x, ","), nil
}

