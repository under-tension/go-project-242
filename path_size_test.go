package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSizeForEmptyDir(t *testing.T) {
	require := require.New(t)

	var a int64 = 0
	b, err := getSize("testdata/empty_dir", false, false)

	require.NoError(err)
	require.Equal(a, b, "Wrong size for empty dir")
}

func TestGetSizeForNotEmptyDir(t *testing.T) {
	require := require.New(t)

	var a int64 = 9
	b, err := getSize("testdata/size_9_byte", false, false)

	require.NoError(err)
	require.Equal(a, b, "Wrong size for empty dir")
}

func TestGetSizeForEmptyFile(t *testing.T) {
	require := require.New(t)

	var a int64 = 0
	b, err := getSize("testdata/empty_file.txt", false, false)

	require.NoError(err)
	require.Equal(a, b, "Wrong size for empty dir")
}

func TestGetSizeForNotEmptyFile(t *testing.T) {
	require := require.New(t)

	var a int64 = 12
	b, err := getSize("testdata/size_12_byte.txt", false, false)

	require.NoError(err)
	require.Equal(a, b, "Wrong size for empty dir")
}
