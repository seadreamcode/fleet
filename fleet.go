package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fleet"
	app.Usage = "Minimal CI System"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello Fleet!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
