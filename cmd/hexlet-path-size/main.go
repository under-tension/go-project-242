package main

import (
	"context"
	"fmt"
	"os"

	cli "github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory;",
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("Hello from Hexlet!")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err.Error())
	}
}
