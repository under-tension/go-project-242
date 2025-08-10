package code

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, false, all)

	if err != nil {
		return "", err
	}

	sizeAndUnit, err := FormatSize(size, human)

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s\t%s", sizeAndUnit, path)

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

			if !all && strings.HasPrefix(entry.Name(), ".") {
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

func FormatSize(size int64, human bool) (string, error) {
	if !human {
		return fmt.Sprintf("%dB", size), nil
	}

	unitNames := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	calcSize := float64(size)
	resultUnitName := "B"
	for _, unitName := range unitNames {
		resultUnitName = unitName

		if calcSize < 1024.0 {
			break
		}

		calcSize /= 1024.0
	}

	if calcSize != math.Trunc(calcSize) {
		return fmt.Sprintf("%.1f%s", calcSize, resultUnitName), nil
	}

	return fmt.Sprintf("%.0f%s", calcSize, resultUnitName), nil
}
