package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var minUnitName = "B"
var unitNames []string = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
var divForNextThreshold float64 = 1024.0

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, recursive, all)

	if err != nil {
		return "", err
	}

	res, err := FormatSize(size, human)

	if err != nil {
		return "", err
	}

	return res, err
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
			if !recursive && entry.IsDir() {
				continue
			}

			if !all && strings.HasPrefix(entry.Name(), ".") {
				continue
			}

			if entry.IsDir() {

				dirSize, err := getSize(filepath.Join(path, entry.Name()), recursive, all)

				if err != nil {
					return 0, err
				}

				size += dirSize
			} else {
				fileSize, err := getSize(filepath.Join(path, entry.Name()), recursive, all)

				if err != nil {
					return 0, err
				}

				size += fileSize
			}
		}
	} else {
		size = f_info.Size()
	}

	return size, nil
}

func FormatSize(size int64, human bool) (string, error) {
	if !human {
		return fmt.Sprintf("%dB", size), nil
	}

	calcSize := float64(size)
	resultUnitName := "B"
	for _, unitName := range unitNames {
		resultUnitName = unitName

		if calcSize < divForNextThreshold {
			break
		}

		calcSize /= divForNextThreshold
	}

	if resultUnitName == minUnitName {
		return fmt.Sprintf("%.0f%s", calcSize, resultUnitName), nil
	}

	return fmt.Sprintf("%.1f%s", calcSize, resultUnitName), nil
}
