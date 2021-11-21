package crypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/url"
	"sync"

	"github.com/ArisAachen/experience/common"
	"github.com/ArisAachen/experience/define"
	"github.com/golang/protobuf/proto"
)

type Cryptor struct {
	// TODO public key and private key can optimize
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey

	lock sync.Mutex
}

// NewCryptor create crypto obj
func NewCryptor() *Cryptor {
	cy := &Cryptor{
	}
	return cy
}

func (cy *Cryptor) GetConfigPath() string {
	return ""
}

// SaveToFile save rsa key to file
// now rsa key is const,
// but in the future, may allow update rsa key from webserver
func (cy *Cryptor) SaveToFile(filename string) error {
	return nil
}

// LoadFromFile load rsa key from config file
func (cy *Cryptor) LoadFromFile(filepath string) error {
	// lock op
	// read rsa keys from config
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		// return err
	}

	// unmarshal buf to key buf
	// rsa key must read successfully
	var key define.RsaKey
	err = proto.Unmarshal(buf, &key)
	if err != nil {
		return err
	}

	// public key and private key must both exist
	if key.Public == "" || key.Private == "" {
		key.Public = defaultPublic
		key.Private = defaultPrivate
		// return errors.New("at least one of rsa keys is nil")
	}

	// parse public key buf to rsa public key
	// dont care about result buf
	block, _ := pem.Decode([]byte(key.Public))
	if block == nil {
		return errors.New("pem decode rsa key is nil")
	}
	// parse block to public key
	cy.publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return err
	}

	// parse private key buf to rsa private key
	// dont care about result buf
	block, _ = pem.Decode([]byte(key.Private))
	if block == nil {
		return errors.New("pem decode rsa key is nil")
	}
	// parse block to private key
	cy.privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	return nil
}

// Encode encode sender to web data
func (cy *Cryptor) Encode(msg string) (define.CryptResult, error) {
	// create crypt result
	var result define.CryptResult
	// check if public key is valid
	if cy.publicKey == nil {
		return result, errors.New("public key has not init yet")
	}
	// check if msg is nil
	if msg == "" {
		return result, errors.New("encode msg is nil")
	}
	rand := common.GetRandomString(16)
	msg = rand + msg
	// use aes-cbc to encode sending data
	// use pkcs7 to padding msg
	padResult := common.PKCSEncode([]byte(msg), define.AecCbcBlockSize)
	// use aes-cbc to encode origin data
	var err error
	result, err = common.AesEncode(padResult)
	if err != nil {
		return result, err
	}
	// use base64 to encode data
	result.Data = base64.StdEncoding.EncodeToString([]byte(result.Data))

	// use rsa to encode sending key
	result.Key, err = common.RSAEncode(cy.publicKey, result.Key)
	if err != nil {
		return result, err
	}
	// use base64 and url to encode key
	result.Key = base64.StdEncoding.EncodeToString([]byte(result.Key))
	result.Key = url.QueryEscape(result.Key)

	return result, nil
}

// Decode decode data from web server
func (cy *Cryptor) Decode(msg define.CryptResult) (string, error) {
	// use rsa to decode key
	// url decode key
	key, err := url.QueryUnescape(msg.Key)
	if err != nil {
		return "", err
	}
	// base64 decode key
	buf, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	// use rsa to decode key
	msg.Key, err = common.RSADecode(cy.privateKey, string(buf))
	if err != nil {
		return "", err
	}
	// use aes-cbc to decode data
	// use base64 to decode data
	buf, err = base64.StdEncoding.DecodeString(msg.Data)
	if err != nil {
		return "", err
	}
	msg.Data = string(buf)
	// use aes-cbc to decode data
	result, err := common.AesDecode(msg)
	if err != nil {
		return "", err
	}
	// use pkcs to decode data
	result, err = common.PKCSDecode(result)
	if err != nil {
		return "", err
	}
	return result, nil
}
