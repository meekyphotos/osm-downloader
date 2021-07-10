package main

import (
	"deploy-cli/pkg/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	downloader := commands.OsmDownloader{}
	err := downloader.Init()
	if err != nil {
		panic(err)
	}

	app := &cli.App{
		Name:        "osm",
		Description: "Download extract CLI",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "path", Value: ".", Usage: "Specify path"},
		},
		Commands: []*cli.Command{
			{
				Name:  "download",
				Usage: "Download a dataset",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "region", Value: "", Usage: "Region to download", Required: true},
					&cli.StringFlag{Name: "format", Value: "pbf", Usage: "Specify format"},
				},
				Action: downloader.OsmDownload,
			},
			{
				Name:  "list",
				Usage: "List available datasets",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "region", Value: "", Usage: "Region to download", Required: true},
					&cli.StringFlag{Name: "format", Value: "pbf", Usage: "Specify format"},
				},
				Action: downloader.List,
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
