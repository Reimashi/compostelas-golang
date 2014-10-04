package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
)

func GenerateRsaKeys(dir string, name string) error {
	// Obtenemos el directorio actual
	rpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	// Generamos y formateamos las llaves
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	privateKeyByteCode := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyByteCode,
	}

	publicKeyByteCode, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	publicKeyBlock := pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyByteCode,
	}

	// Creamos y escribimos los ficheros
	filepublic, err := os.Create(rpath + string(os.PathSeparator) + name + ".public")
	if err != nil {
		return err
	}

	defer filepublic.Close()

	_, err = filepublic.Write([]byte(pem.EncodeToMemory(&publicKeyBlock)))
	if err != nil {
		return err
	}

	filepublic.Sync()

	fileprivate, err := os.Create(rpath + string(os.PathSeparator) + name + ".private")
	if err != nil {
		return err
	}

	defer fileprivate.Close()

	_, err = fileprivate.Write([]byte(pem.EncodeToMemory(&privateKeyBlock)))
	if err != nil {
		return err
	}

	fileprivate.Sync()

	return nil
}
