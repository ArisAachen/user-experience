package common

import (
	"math/rand"
	"time"
)

// randPool offer pool to generate random aes-cbc key
var randPool string

// random use 'num char _' to create random 32 byte aes-cbc key
func random(length int) string {
	// check if size is valid
	if length <= 0 {
		return ""
	}
	// create random string
	var result string
	// get random seed
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	// check randPool size
	size := len(randPool)
	if len(randPool) == 0 {
		return ""
	}
	// random each ch
	for index := 0; index < length; index++ {
		result += string(randPool[seed.Intn(size)])
	}
	return result
}

func init() {
	randPool = "123456789_"
	// add all character to pool
	for ch := 'a'; ch < 'Z'; ch++ {
		randPool += string(ch)
	}
}
