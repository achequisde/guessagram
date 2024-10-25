package helpers

import "math/rand"

// Shuffles array using the Fisher-Yates algorithm
func Shuffle[K interface{}](items []K) []K {
	result := make([]K, len(items))
	copy(result, items)

	for i := len(result) - 1; i > 0; i-- {
		j := rand.Intn(i)

		tmp := result[j]
		result[j] = result[i]
		result[i] = tmp
	}

	return result
}
