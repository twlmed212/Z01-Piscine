package piscine

func SplitWhiteSpaces(s string) []string {
	var saveString []string
	var saveWord string
	for i, getSpaces := range s {

		if !(getSpaces == ' ' || getSpaces == '\n' || getSpaces == '\t') {
			saveWord += string(getSpaces)
		}

		// if getSpaces == ' ' || getSpaces == '9' || getSpaces == '\n' {
		if getSpaces == ' ' || i == len(s)-1 {
			if len(saveWord) > 0 {
				// fmt.Println(len(s) - 1)
				saveString = append(saveString, saveWord)
			}
			saveWord = ""
		}

	}
	return saveString
}
