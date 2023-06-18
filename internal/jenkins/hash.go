package jenkins

func Hash(input string) uint32 {
	hash := uint32(0)

	for _, char := range input {
		hash += uint32(char)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash
}
