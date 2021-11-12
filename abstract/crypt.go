package abstract

// Encoder use to encrypt data
// now has two type encoder: aes-cbc rsa
type Encoder interface {
	Encode(msg string)
}

// Decoder use to decrypt data
// now has two type decoder: aes-cbc rsa
type Decoder interface {
	Decode(msg string)
}
