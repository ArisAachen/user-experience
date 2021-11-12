package common

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ArisAachen/experience/define"
)

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

// pkcsEncode use to padding pkcs data
// now pkcs encode use to encode pkcs7 and pkcs5(const blockSize 8)
func pkcsEncode(msg string, blockSize int) string {
	// cal padding size
	size := len(msg)
	padSize := blockSize - (size % blockSize)
	if padSize == 0 {
		padSize = blockSize
	}
	// get repeat
	padMsg := strconv.Itoa(padSize)
	// create padding result
	result := msg + strings.Repeat(padMsg, padSize)
	return result
}

// PKCSDecode use to unPadding pkcs data
// now pkcs decode use to decode pkcs7 and pkcs5(const blockSize 8)
func pkcsDecode(msg string) (result string, err error) {
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
func AesEncode(msg string) (define.AesResult, error) {
	// in case slice out of bounder
	// this will not happens, but maintainer make stupid mistake
	defer func() {
	}()
	// create 32 byte aes-cbc key
	var result define.AesResult
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
	mode.CryptBlocks([]byte(result.Data), []byte(msg))
	return result, nil
}

// AesDecode use aes-cbc decode msg
// after aes-cbc decode, also should use un-padding to parse data
func AesDecode(msg define.AesResult) (string, error) {
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
	var result string
	mode.CryptBlocks([]byte(result), []byte(msg.Data))
	return result, nil
}
