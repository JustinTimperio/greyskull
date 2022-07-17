package main

import (
	"errors"
	"log"
	"os"

	"github.com/cloudflare/circl/pke/kyber/kyber1024"
	cli "github.com/urfave/cli/v2"
)

func main() {
	var (
		seed    string
		pubKey  *kyber1024.PublicKey
		privKey *kyber1024.PrivateKey
	)

	homepath, _ := os.UserHomeDir()

	app := &cli.App{
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "genKeys",
				Usage: "Generates a public and private Kyber key. If the seed flag is not present, crypto/rand will seed the keys.",
			},
			&cli.StringFlag{
				Name:        "keySeed",
				Usage:       "Sets a manual deterministic seed for creating a public/private key pair.",
				Destination: &seed,
			},
			&cli.StringFlag{
				Name:  "keyPath",
				Usage: "Sets a path to the kyber keyset. Defaults to ~/.greyskull for each user.",
				Value: ".greyskull",
			},
		},
		Action: func(cCtx *cli.Context) (err error) {
			if !pathExists(homepath + "/" + cCtx.String("keyPath")) {
				err := os.Mkdir(homepath+"/"+cCtx.String("keyPath"), 0755)
				if err != nil {
					return err
				}
			}

			if cCtx.Bool("genKeys") {
				pubPath := homepath + "/" + cCtx.String("keyPath") + "/kyber.pub"
				privPath := homepath + "/" + cCtx.String("keyPath") + "/kyber.priv"
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
					log.Print("Gen Seed Key")
					// NewKeyFromSeed will panic if seed length is not == KeySeedSize
					if len(seed) != int(kyber1024.KeySeedSize) {
						return errors.New("Key seed length is not 32 chars")
					}
					pubKey, privKey, err = GenKeys([]byte(seed))
				} else {
					pubKey, privKey, err = GenKeys(nil)
					log.Print("Gen Rand Key")
				}
				if err != nil {
					return err
				}

				err = StoreKeys(pubKey, privKey, pubPath, privPath)
				if err != nil {
					return err
				}
			}

			return err
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
