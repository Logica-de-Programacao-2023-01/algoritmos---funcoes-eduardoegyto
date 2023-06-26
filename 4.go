func segundo(x []int) (int, error) {
	if len(x) < 2 {
		return 0, errors.New("O slice precisa ter pelo menos 2 elementos")
	}
	max := x[0]
	y := numbers[1]
	if y > max {
		max, y = y, max
	}
	for _, num := range x[2:] {
		if num > max {
			y = max
			max = num
		} else if num > y {
			y = num
		}
	}
	return y, nil
}
