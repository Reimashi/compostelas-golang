package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

var default_iv string = "22277966616d9bc47177bd02603d08c9a67d5380d0fe8cf3b44438dff7b9"

func EncryptAes(data []byte, key string) ([]byte, error) {
	encrypted := make([]byte, len(data))

	iv, err := hex.DecodeString(default_iv)
	if err != nil {
		return encrypted, err
	}

	aesBlockEncrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return encrypted, err
	}

	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, data)

	return encrypted, nil
}

func DecryptAes(data []byte, key string) ([]byte, error) {
	decrypted := make([]byte, len(data))

	iv, err := hex.DecodeString(default_iv)
	if err != nil {
		return decrypted, err
	}

	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return decrypted, err
	}

	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, data)

	return decrypted, nil
}
