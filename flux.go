package main

import (
	"flux/app/cmd"

	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Flux"
	app.Usage = "A self-hosted Minecraft modpack service."
	app.Authors = []*cli.Author{
		{
			Name:  "Shane Clark",
			Email: "sclark@frostfire.io",
		},
	}
	app.Version = "0.0.1+dev"
	app.Commands = []*cli.Command{
		&cmd.Serve,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to start applicaiton: %v", err)
	}
}
