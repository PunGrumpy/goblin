package jenkins

import "github.com/PunGrumpy/goblin/external/character"

func FindPreimages(target uint32, length int, characters []rune) []string {
	minChar, maxChar := character.GetCharacterBounds(characters)

	hash := target
	hash *= 0x3FFF8001
	hash ^= (hash >> 11) ^ (hash >> 22)
	hash *= 0x38E38E39

	output := make([]string, 0)

	if length > 5 {
		outputChan := make(chan string)
		defer close(outputChan)

		for c := minChar; c <= maxChar; c++ {
			go reverse(hash, make([]rune, length), length-1, characters, minChar, maxChar, rune(c), outputChan)
		}

		for partialOutput := range outputChan {
			output = append([]string{partialOutput}, output...)
		}
	} else {
		reverse(hash, make([]rune, length), length-1, characters, minChar, maxChar, 0, make(chan string, 1) /* outputChan */)
	}

	return output
}

func reverse(hash uint32, buffer []rune, depth int, characters []rune, minChar, maxChar rune, forceChar rune, outputChan chan<- string) {
	hash ^= (hash >> 6) ^ (hash >> 12) ^ (hash >> 18) ^ (hash >> 24) ^ (hash >> 30)
	hash *= 0xC00FFC01

	if depth == 0 {
		if hash < uint32(minChar) || hash > uint32(maxChar) {
			return
		}

		buffer[0] = rune(hash)

		outputChan <- string(buffer)

		return
	}

	recur := func(ch rune) {
		buffer[depth] = ch
		reverse((hash - uint32(ch)), buffer, depth-1, characters, minChar, maxChar, 0, outputChan)
	}

	if forceChar != 0 {
		recur(forceChar)
	} else {
		for _, ch := range characters {
			recur(ch)
		}
	}
}
