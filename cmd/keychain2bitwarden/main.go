package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ItsOnlyGame/keychain2bitwarden/internal/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "keychain2bitwarden",
		Usage: "Converts Apple Keychain files to Bitwarden JSON files",
		Commands: []*cli.Command{
			{
				Name:  "convert",
				Usage: "Converts (reads and exports) a file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "import-path",
						Usage:    "The file to convert",
						Aliases:  []string{"i"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "export-path",
						Usage:    "Filename to export to",
						Aliases:  []string{"e"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					err := utils.ConvertKeychainToJSON(c.String("import-path"), c.String("export-path"))

					if err != nil {
						fmt.Println("Covenrsion failed")
						return err
					}
					fmt.Println("Successfully converted file")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
