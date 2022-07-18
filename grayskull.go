package main

import (
	"errors"
	// "fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	var (
		seed        string
		filePath    string
		keysetPath  string
		homepath, _ = os.UserHomeDir()
	)

	app := &cli.App{
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			// Keyset Flags
			&cli.BoolFlag{
				Name:  "genKeyset",
				Usage: "Generates a public and private Kyber key. If the seed flag is not present, crypto/rand will seed the keys.",
			},
			&cli.StringFlag{
				Name:        "keysetSeed",
				Usage:       "Sets a deterministic seed for creating a public/private key pair.",
				Destination: &seed,
			},
			&cli.StringFlag{
				Name:  "keysetPath",
				Usage: "Sets a path to the kyber keyset. Defaults to ~/.greyskull for each user.",
				Value: ".greyskull",
			},

			// Encrypt/Decrypt Flags
			&cli.BoolFlag{
				Name:  "encrypt",
				Usage: "Encrypt a file using Kyber 1024",
			},
			&cli.BoolFlag{
				Name:  "decrypt",
				Usage: "Decrypt a file using Kyber 1024",
			},
			&cli.StringFlag{
				Name:        "filePath",
				Usage:       "Path to the file that you want to try to encrypt or decrypt.",
				Destination: &filePath,
			},
			&cli.StringFlag{
				Name:        "pubKeyPath",
				Usage:       "Path to the publicKey that you want to use to encrypt the file.",
				Destination: &keysetPath,
			},
		},

		// Cli Main
		Action: func(cCtx *cli.Context) (err error) {
			// Make sure that the greyskull key directory exists
			if !pathExists(homepath + "/" + cCtx.String("keysetPath")) {
				err := os.Mkdir(homepath+"/"+cCtx.String("keysetPath"), 0755)
				if err != nil {
					return err
				}
			}

			// Generate Kyber keys
			if cCtx.Bool("genKeyset") {
				err = CreateKyberKeys(cCtx, homepath, seed)
				if err != nil {
					return err
				}
			}

			// Encrypt using public key
			if cCtx.Bool("encrypt") {
				if cCtx.String("pubKeyPath") == "" {
					return errors.New("No path was given to a public key")
				}
				pk, err := ReadPublicKey(cCtx.String("pubKeyPath"))
				if err != nil {
					return err
				}
				//
				if cCtx.String("filePath") == "" {
					return errors.New("No path was given to a file to encrypt")
				}
				ePath, err := EncryptFile(pk, keysetPath)
				if err != nil {
					return err
				}
				log.Print("Encrypted file at " + ePath)
			}

			// Decrypt using public key
			if cCtx.Bool("decrypt") {
				if cCtx.String("filePath") == "" {
					return errors.New("No path was given to a file to decrypt")
				}

				pk, err := ReadPrivateKey(cCtx.String("filePath"))
				if err != nil {
					return err
				}

				dPath, err := DecryptFile(pk, filePath)
				if err != nil {
					return err
				}
				log.Print("Decrypted file at " + dPath)
			}

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
