package util

import (
	"math/rand"
	"time"
)

// Rand generates a random number
func Rand(n int) int {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return seed.Intn(n)
}
