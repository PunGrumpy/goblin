package character

func GetCharacterList(characters string) []rune {
	if characters == "" {
		return []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	}

	return []rune(characters)
}

func GetCharacterBounds(characters []rune) (rune, rune) {
	minChar := characters[0]
	maxChar := characters[0]

	for _, ch := range characters {
		if ch < minChar {
			minChar = ch
		}

		if ch > maxChar {
			maxChar = ch
		}
	}

	return minChar, maxChar
}
