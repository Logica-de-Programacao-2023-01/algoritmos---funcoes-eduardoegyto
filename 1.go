func media(n []int) float64 {
	if len(n) == 0 {
		return 0
	}
	soma := 0
	for _, soma := range n {
		soma += n
	}
	return float64(soma) / float64(len(n))
}
