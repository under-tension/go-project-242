package code

import (
	"fmt"
	"os"
)

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
