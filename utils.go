package main

import (
	"crypto/rand"

	"github.com/cloudflare/circl/pke/kyber/kyber1024"
)

// GenKeys Gen Post Quantum Public and Private Keys
func GenKeys(seed []byte) (pubKey *kyber1024.PublicKey, privKey *kyber1024.PrivateKey, err error) {
	// Get cryptographically secure rand stream from getrandom(2)
	// TODO: Upgrade this in the future
	r := rand.Reader

	if seed == nil {
		// Gen Kyber Keys
		pubKey, privKey, err = kyber1024.GenerateKey(r)
	} else {
		pubKey, privKey = kyber1024.NewKeyFromSeed(seed)
	}
	return pubKey, privKey, err
}

// StoreKeys writes a private and public key to ~/.greyskull
func StoreKeys(pubKey *kyber1024.PublicKey, privKey *kyber1024.PrivateKey, pubPath string, privPath string) (err error) {
	var (
		pubk  = make([]byte, int(kyber1024.PublicKeySize))
		privk = make([]byte, int(kyber1024.PrivateKeySize))
	)

	pubKey.Pack(pubk)
	err = createFile(pubPath, pubk, 0755)
	if err != nil {
		return err
	}

	privKey.Pack(privk)
	err = createFile(privPath, privk, 0644)
	if err != nil {
		return err
	}

	return nil
}
