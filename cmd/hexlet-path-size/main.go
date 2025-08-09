package main

import (
	"context"
	"errors"
	"fmt"
	"os"

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

			resStr, err := GetPathSize(path, false, true, false)

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

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, false, false)

	result := fmt.Sprintf("%d\t%s", size, path)

	return result, err
}

func getSize(path string, recursive, all bool) (int64, error) {
	f_info, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}

	var size int64 = 0

	if f_info.IsDir() {
		d_entry, err := os.ReadDir(path)

		if err != nil {
			return 0, err
		}

		for _, entry := range d_entry {
			if entry.IsDir() {
				continue
			}

			f_info, err := entry.Info()

			if err != nil {
				return 0, err
			}

			size += f_info.Size()
		}
	} else {
		size = f_info.Size()
	}

	return size, nil
}
