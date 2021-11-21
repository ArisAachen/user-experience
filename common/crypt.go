package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/ArisAachen/experience/define"
)

// PKCSEncode use to padding pkcs data
// now pkcs encode use to encode pkcs7 and pkcs5(const blockSize 8)
func PKCSEncode(msg []byte, blockSize int) string {
	// cal padding size
	size := len(msg)
	padSize := blockSize - (size % blockSize)
	if padSize == 0 {
		padSize = blockSize
	}
	// get repeat
	repeat := bytes.Repeat([]byte(string(rune(padSize))), padSize)
	// create padding result
	result := append(msg, repeat...)
	return string(result)
}

// PKCSDecode use to unPadding pkcs data
// now pkcs decode use to decode pkcs7 and pkcs5(const blockSize 8)
func PKCSDecode(msg string) (result string, err error) {
	// in case slice out of bounder
	defer func() {
		// try to catch panic
		cover := recover()
		if cover == nil {
			return
		}
		// format error
		err = fmt.Errorf("%v", cover)
		return
	}()
	// cal size
	size := len(msg)
	// check size
	if size == 0 {
		err = errors.New("pkcs7 decode empty msg")
		return
	}
	// get padding num
	padSize := int(msg[size-1])
	result = msg[:size-padSize]
	return
}

// AesEncode use aes-cbc encode msg,
// also aes-cbc should call after padding
// aes cbc key is random created every time
func AesEncode(msg string) (define.CryptResult, error) {
	// in case slice out of bounder
	// this will not happens, but maintainer make stupid mistake
	defer func() {
	}()
	// create 32 byte aes-cbc key
	var result define.CryptResult
	key := random(define.AesCbcKeySize)
	iv := key[:define.AesCbcIvVec]
	// save key
	result.Key = key
	// create cipher obj
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return result, err
	}
	// create block mode
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	// use pkcs7 to padding
	// in case crypt mode panic, should check size here
	if len(msg)%mode.BlockSize() != 0 {
		return result, errors.New("aes encrypt not full block")
	}
	// use aes-cbc encrypt
	buf := make([]byte, len([]byte(msg)))
	mode.CryptBlocks(buf, []byte(msg))
	result.Data = string(buf)
	return result, nil
}

// AesDecode use aes-cbc decode msg
// after aes-cbc decode, also should use un-padding to parse data
func AesDecode(msg define.CryptResult) (string, error) {
	// in case slice out of bounder
	// this will not happens, but maintainer make stupid mistake
	defer func() {
	}()
	// check if decode aes-cbc data is valid, now aes-key must be 32 size
	if msg.Key == "" || msg.Data == "" || len(msg.Key) != define.AesCbcKeySize {
		return "", errors.New("aes decode invalid param")
	}
	// create cipher obj
	block, err := aes.NewCipher([]byte(msg.Key))
	if err != nil {
		return "", err
	}
	// create decrypt block mode
	iv := msg.Key[:define.AesCbcIvVec]
	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	// in case crypt mode panic, should check size here
	if len(msg.Data)%mode.BlockSize() != 0 {
		return "", errors.New("aes decrypt not full block")
	}
	buf := make([]byte, len([]byte(msg.Data)))
	mode.CryptBlocks(buf, []byte(msg.Data))
	return string(buf), nil
}

// RSAEncode use rsa to encode data
func RSAEncode(key *rsa.PublicKey, msg string) (string, error) {
	// check if public key and message is valid
	if key == nil || msg == "" {
		return "", errors.New("rsa encode param is invalid")
	}
	// use rsa to encode msg
	result, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(msg))
	if err != nil {
		return "", err
	}
	// encode rsa message success
	return string(result), nil
}

// RSADecode use rsa to decode data
func RSADecode(key *rsa.PrivateKey, msg string) (string, error) {
	// check if private key and message is valid
	if key == nil || msg == "" {
		return "", errors.New("rsa encode param is invalid")
	}
	// use rsa to decode msg
	result, err := rsa.DecryptPKCS1v15(rand.Reader, key, []byte(msg))
	if err != nil {
		return "", err
	}
	// decode rsa message success
	return string(result), nil
}

// GetRandomString get random string
func GetRandomString(length int) string {
	return random(length)
}
