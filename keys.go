package main

import (
	"crypto/rand"
	"errors"

	"github.com/cloudflare/circl/pke/kyber/kyber1024"
	cli "github.com/urfave/cli/v2"
)

// genKeys Gen Post Quantum Public and Private Keys
func genKeys(seed []byte) (pubKey *kyber1024.PublicKey, privKey *kyber1024.PrivateKey, err error) {
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

// storeKeys writes a private and public key to ~/.greyskull
func storeKeys(pubKey *kyber1024.PublicKey, privKey *kyber1024.PrivateKey, pubPath string, privPath string) (err error) {
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

// CreateKyberKeys creates a set of Kyber keys and writes them to the disk
func CreateKyberKeys(cCtx *cli.Context, homepath string, seed string) (err error) {
	var (
		pubKey   *kyber1024.PublicKey
		privKey  *kyber1024.PrivateKey
		pubPath  = homepath + "/" + cCtx.String("keyPath") + "/kyber.pub"
		privPath = homepath + "/" + cCtx.String("keyPath") + "/kyber.priv"
	)

	// Make sure user doesn't overwrite their keys
	if pathExists(pubPath) {
		if !askForConfirmation("Do you want to overwrite your existing public key?") {
			return errors.New("User aborted key creation")
		}
	}
	if pathExists(privPath) {
		if !askForConfirmation("Do you want to overwrite your existing private key?") {
			return errors.New("User aborted key creation")
		}
	}

	if seed != "" {
		// NewKeyFromSeed will panic if seed length is not == KeySeedSize
		if len(seed) != int(kyber1024.KeySeedSize) {
			return errors.New("Key seed length is not 32 chars")
		}
		pubKey, privKey, err = genKeys([]byte(seed))
	} else {
		pubKey, privKey, err = genKeys(nil)
	}
	if err != nil {
		return err
	}

	err = storeKeys(pubKey, privKey, pubPath, privPath)
	if err != nil {
		return err
	}
	return nil
}
