package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	var (
		seed string
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
			// Make sure that the greyskull key directory exists
			if !pathExists(homepath + "/" + cCtx.String("keyPath")) {
				err := os.Mkdir(homepath+"/"+cCtx.String("keyPath"), 0755)
				if err != nil {
					return err
				}
			}

			if cCtx.Bool("genKeys") {
				err = CreateKyberKeys(cCtx, homepath, seed)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
