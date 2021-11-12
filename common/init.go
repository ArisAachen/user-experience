package common

// randPool offer pool to generate random aes-cbc key
var randPool string

func init() {
	randPool = "123456789_"
	// add all character to pool
	for ch := 'a'; ch < 'Z'; ch++ {
		randPool += string(ch)
	}
}
