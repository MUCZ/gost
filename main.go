package main

import (
	"errors"
	"gost/client"
	"gost/server"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

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
					if cCtx.NArg() != 1 {
						color.Red("uid is required")
						return nil
					}
					uid := cCtx.Args().Get(0)
					color.Blue("getting uid: %s", uid)
					ret, err := client.Get(uid)
					if err != nil {
						color.Red("error: %s", err)
					} else {
						color.Green("ret: ")
						color.Cyan(ret)
					}
					return nil
				},
			},
			{
				Name:    "set",
				Aliases: []string{"st"},
				Usage:   "set a gist from a file, show the uid",
				Action: func(cCtx *cli.Context) error {
					if cCtx.NArg() != 1 {
						return errors.New("file is required")
					}
					file := cCtx.Args().Get(0)
					color.Blue("setting file: %s", file)
					fhandler, err := os.Open(file)
					if err != nil {
						return err
					}
					defer fhandler.Close()
					msg := make([]byte, 20*1024)
					n, err := fhandler.Read(msg)
					if err != nil {
						return err
					}
					if n == 20*1024 {
						return errors.New("msg is too long")
					}
					ret, err := client.Post(string(msg[:n]))
					if err != nil {
						return errors.New("error: " + err.Error() + "ret : " + ret)
					} else {
						color.Green("ret: ")
						color.Cyan(ret)
					}
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"r"},
				Usage:   "delete a gist from uid",
				Action: func(cCtx *cli.Context) error {
					if cCtx.NArg() != 1 {
						return errors.New("uid is required")
					}
					uid := cCtx.Args().Get(0)
					color.Blue("deleting uid: %s", uid)
					ret, err := client.Delete(uid)
					if err != nil {
						return err
					} else {
						color.Green("ret:  ")
						color.Cyan(ret)
					}
					return nil
				},
			},
			{
				Name:    "describe",
				Aliases: []string{"d"},
				Usage:   "describe a gist from uid",
				Action: func(cCtx *cli.Context) error {
					if cCtx.NArg() != 1 {
						return errors.New("uid is required")
					}
					uid := cCtx.Args().Get(0)
					color.Blue("describing uid: %s", uid)
					ret, err := client.Describe(uid)
					if err != nil {
						return err
					} else {
						color.Green("ret: ")
						color.Cyan(ret)
					}
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list all gists",
				Action: func(cCtx *cli.Context) error {
					color.Blue("listing all gists")
					ret, err := client.List()
					if err != nil {
						return err
					} else {
						color.Green("ret: ")
						color.Cyan(ret)
					}
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
					ret, err := client.Check()
					if err != nil {
						return err
					} else {
						color.Green("server is running: ret: %v \n ", ret)
					}
					return nil
				},
			},
		},
	}

	color.Cyan("Server addr: %s", server.Addr)
	if err := app.Run(os.Args); err != nil {
		color.Red(err.Error())
	}

}
