package main

import (
	"errors"

	// "github.com/cloudflare/circl/pke/kyber/kyber1024"
	kyber "github.com/symbolicsoft/kyber-k2so"
	cli "github.com/urfave/cli/v2"
)

// genKeys Gen Post Quantum Public and Private Keys
func genKeys(seed []byte) (pubKey [1568]byte, privKey [3168]byte, err error) {
	privKey, pubKey, err = kyber.KemKeypair1024()
	return pubKey, privKey, err
}

// storeKeys writes a private and public key to ~/.greyskull
func storeKeys(pubKey [1568]byte, privKey [3168]byte, pubPath string, privPath string) (err error) {

	err = createFile(pubPath, pubKey[:], 0755)
	if err != nil {
		return err
	}

	err = createFile(privPath, privKey[:], 0644)
	if err != nil {
		return err
	}

	return nil
}

// CreateKyberKeys creates a set of Kyber keys and writes them to the disk
func CreateKyberKeys(cCtx *cli.Context, homepath string, seed string) (err error) {
	var (
		pubPath  = homepath + "/" + cCtx.String("keysetPath") + "/kyber.pub"
		privPath = homepath + "/" + cCtx.String("keysetPath") + "/kyber.priv"
	)

	// Make sure user doesn't overwrite their keys
	if pathExists(pubPath) {
		if !askForConfirmation("Do you want to overwrite your existing public key at " + pubPath + "?") {
			return errors.New("User aborted key creation")
		}
	}
	if pathExists(privPath) {
		if !askForConfirmation("Do you want to overwrite your existing private key at " + privPath + "?") {
			return errors.New("User aborted key creation")
		}
	}

	pubKey, privKey, err := genKeys(nil)
	if err != nil {
		return err
	}

	err = storeKeys(pubKey, privKey, pubPath, privPath)
	if err != nil {
		return err
	}
	return nil
}
