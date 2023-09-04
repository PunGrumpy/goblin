package jenkins

import (
	"sync"

	"github.com/PunGrumpy/goblin/external/character"
)

const (
	hashMultiplier1 = 0x3FFF8001
	hashMultiplier2 = 0x38E38E39
	hashXOR1        = 11
	hashXOR2        = 22
)

func FindPreimages(target uint32, length int, characters []rune) []string {
	minChar, maxChar := character.GetCharacterBounds(characters)

	hash := target
	hash *= hashMultiplier1
	hash ^= (hash >> hashXOR1) ^ (hash >> hashXOR2)
	hash *= hashMultiplier2

	output := make([]string, 0)

	if length > 5 {
		var wg sync.WaitGroup
		outputChan := make(chan string, len(characters))

		for _, ch := range characters {
			wg.Add(1)
			go func(ch rune) {
				defer wg.Done()
				reverse(hash-uint32(ch), make([]rune, length), length-1, characters, minChar, maxChar, ch, outputChan)
			}(ch)
		}

		go func() {
			wg.Wait()
			close(outputChan)
		}()

		for partialOutput := range outputChan {
			output = append(output, partialOutput)
		}
	} else {
		reverse(hash, make([]rune, length), length-1, characters, minChar, maxChar, 0, output)
	}

	return output
}

func reverse(hash uint32, buffer []rune, depth int, characters []rune, minChar, maxChar rune, forceChar rune, output interface{}) {
	hash ^= (hash >> 6) ^ (hash >> 12) ^ (hash >> 18) ^ (hash >> 24) ^ (hash >> 30)
	hash *= 0xC00FFC01

	if depth == 0 {
		if hash < uint32(minChar) || hash > uint32(maxChar) {
			return
		}

		buffer[0] = rune(hash)

		if outputChan, ok := output.(chan string); ok {
			outputChan <- string(buffer)
		} else if outputSlice, ok := output.([]string); ok {
			_ = append(outputSlice, string(buffer))
		}
		return
	}

	recur := func(ch rune) {
		buffer[depth] = ch
		reverse(hash-uint32(ch), buffer, depth-1, characters, minChar, maxChar, 0, output)
	}

	if forceChar != 0 {
		recur(forceChar)
	} else {
		for _, ch := range characters {
			recur(ch)
		}
	}
}
