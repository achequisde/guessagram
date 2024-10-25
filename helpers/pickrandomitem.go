package helpers

import "math/rand"

func PickRandomItem[K interface{}](items []K) K {
	return items[rand.Intn(len(items))]
}
