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
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory;",
		Action: func(context.Context, *cli.Command) error {
			if len(os.Args) < 2 {
				return errors.New("file path not found")
			}

			path := os.Args[1]

			resStr, err := code.GetPathSize(path, false, true, false)

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
