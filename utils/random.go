package utils

import (
	"math/rand"
	"time"
)

// global seed for random generation of numbers
var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomChoice returns a random selection from a list of strings
func RandomChoice(choices []string) string {
	n := seed.Int() % len(choices)
	return choices[n]
}

// RandomBool returns a random boolean choice
func RandomBool() bool {
	return rand.Float32() < 0.5
}

func RandomInt(maxval int) int {
	return seed.Intn(maxval)
}

func RandomUint64() uint64 {
	return seed.Uint64()
}

// RandomBool returns a boolean choice at the user's threshold
func RandomBoolWeight(chanceTrue float32) bool {
	return rand.Float32() < chanceTrue
}

// characters to choose from
const charset = "abcdefghijklmnopqrstuvwxyz"
const capitals = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Get a random character
func RandomChar() string {
	chars := charset + capitals
	return string(chars[seed.Intn(len(chars))])
}

func RandomRange(min int, max int) int {
	return seed.Intn(max-min) + min
}

// RandomName generates a random name for a function, variaable, etc.
func RandomName() string {
	length := seed.Intn(20-10) + 10
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seed.Intn(len(charset))]
	}
	return string(result)
}
