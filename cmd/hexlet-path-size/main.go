package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"code"

	cli "github.com/urfave/cli/v3"
)

func main() {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "human",
			Usage:   "human-readable sizes (auto-select unit)",
			Value:   false,
			Aliases: []string{"H"},
		},
		&cli.BoolFlag{
			Name:    "all",
			Usage:   "include hidden files and directories",
			Value:   false,
			Aliases: []string{"a"},
		},
		&cli.BoolFlag{
			Name:    "recursive",
			Usage:   "recursive size of directories",
			Value:   false,
			Aliases: []string{"r"},
		},
	}

	cmd := &cli.Command{
		Name:            "hexlet-path-size",
		Usage:           "print size of a file or directory;",
		Flags:           flags,
		HideHelpCommand: true,
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "path",
				Value:     "",
				UsageText: "path to file or directory",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.StringArg("path")
			human := cmd.Bool("human")
			all := cmd.Bool("all")
			recursive := cmd.Bool("recursive")

			if path == "" {
				return errors.New("file path not found")
			}
			fmt.Println(recursive)

			resStr, err := code.GetPathSize(path, recursive, human, all)

			if err != nil {
				return err
			}

			fmt.Println(resStr)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err.Error())
	}
}
