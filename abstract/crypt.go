package abstract

import "github.com/ArisAachen/experience/define"

// Encoder use to encrypt data
// now has two type encoder: aes-cbc rsa
type Encoder interface {
	Encode(msg string) (define.CryptResult, error)
}

// Decoder use to decrypt data
// now has two type decoder: aes-cbc rsa
type Decoder interface {
	Decode(msg define.CryptResult) (string, error)
}
