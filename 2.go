func vogais(x string) int {
	vogal := "aeiouAEIOU"
	conta := 0
	for _, qnt := range x {
		if strings.ContainsRune(vogal, qnt) {
			conta++
		}
	}
	return conta
}
