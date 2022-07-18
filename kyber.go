package main

import (
	// "fmt"
	kyber "github.com/symbolicsoft/kyber-k2so"
)

func EncryptFile(pubKey [1568]byte, filePath string) (encryptedFilePath string, err error) {
	_, _, _ = kyber.KemEncrypt1024(pubKey)
	return "", nil
}

func DecryptFile(privKey [3168]byte, filePath string) (decryptedFilePath string, err error) {
	return "", nil
}

func ReadPrivateKey(filePath string) (privKey [3168]byte, err error) {
	x, err := readFile(filePath)
	if err != nil {
		return privKey, err
	}
	copy(privKey[:], x)
	return privKey, nil
}

func ReadPublicKey(filePath string) (pubKey [1568]byte, err error) {
	x, err := readFile(filePath)
	if err != nil {
		return pubKey, err
	}
	copy(pubKey[:], x)
	return pubKey, nil
}
