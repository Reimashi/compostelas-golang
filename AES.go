package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func encriptar_aes(data []byte, key []byte, iv []byte) []byte {
	encrypted := make([]byte, len(data))
	aesBlockEncrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, data)
	return encrypted
}

func desencriptar_aes(data []byte, key []byte, iv []byte) []byte {
	decrypted := make([]byte, len(data))
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, data)
	return decrypted
}
