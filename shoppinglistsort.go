package piscine

func ShoppingListSort(slice []string) []string {
	for i, getChar := range slice {
		for j, getLenght := range slice {
			if len(getChar) < len(getLenght) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
