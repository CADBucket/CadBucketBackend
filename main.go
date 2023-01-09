package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "CADBucket",
		Usage: "Collaborate On Cad Models",
		Action: func(*cli.Context) error {
			fmt.Println("Hello World")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("err")
	}

}

