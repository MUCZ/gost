package main

import (
	"fmt"
	"gost/client"
	"gost/server"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// get [uid]
// set [filename]
// delete [uid]
// describe [uid]

// start
// serverStatus

func main() {

	app := &cli.App{
		Name:  "gost",
		Usage: "dummy gist written in go",
		Commands: []*cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "get a gist from uid",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					client.Get()
					return nil
				},
			},
			{
				Name:    "set",
				Aliases: []string{"st"},
				Usage:   "set a gist from a file, show the uid",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					client.Post()
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"r"},
				Usage:   "delete a gist from uid",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("deleted task: ", cCtx.Args().First())
					client.Delete()
					return nil
				},
			},
			{
				Name:    "describe",
				Aliases: []string{"d"},
				Usage:   "describe a gist from uid",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("described task: ", cCtx.Args().First())
					client.Describe()
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list all gists",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("listed all gists")
					client.List()
					return nil
				},
			},
			{
				Name:    "start",
				Usage:   "start the server",
				Aliases: []string{"s"},
				Action: func(cCtx *cli.Context) error {
					server.Start()
					return nil
				},
			},
			{
				Name:    "serverStatus",
				Usage:   "get the server status",
				Aliases: []string{"ss"},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("server is running")
					client.Check()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
