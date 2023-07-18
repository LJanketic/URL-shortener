package utils

import (
	"math/rand"
)

var runesList = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// TODO: Alter random with a seed using time.Now()
func RandomizeURL(size int) string {
	randStr := make([]rune, size)

	for i := range randStr {
		randStr[i] = runesList[rand.Intn(len(runesList))]
	}

	return string(randStr)
}
