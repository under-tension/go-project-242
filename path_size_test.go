package code

import (
	"math"
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

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		human    bool
		expected string
	}{
		// --- human = false ---
		{"Non-human: zero", 0, false, "0B"},
		{"Non-human: small", 42, false, "42B"},
		{"Non-human: large", 999999, false, "999999B"},

		// --- human = true: bytes ---
		{"Human: 0B", 0, true, "0B"},
		{"Human: 1B", 1, true, "1B"},
		{"Human: 512B", 512, true, "512B"},
		{"Human: 1023B", 1023, true, "1023B"},

		// --- KB ---
		{"Human: 1024B = 1KB", 1024, true, "1KB"},
		{"Human: 1536B = 1.5KB", 1536, true, "1.5KB"},
		{"Human: 2048B = 2KB", 2048, true, "2KB"},
		{"Human: 1025B ~ 1.0KB", 1025, true, "1.0KB"},
		{"Human: 1100B ~ 1.1KB", 1100, true, "1.1KB"},

		// --- MB ---
		{"Human: 1MB exact", 1024 * 1024, true, "1MB"},
		{"Human: 1.5MB", 3 * 1024 * 512, true, "1.5MB"}, // 1.5 * 1024^2
		{"Human: 2.3MB", int64(math.Round(2.3 * 1024 * 1024)), true, "2.3MB"},

		// --- GB ---
		{"Human: 1GB exact", 1024 * 1024 * 1024, true, "1GB"},
		{"Human: 1.7GB", int64(math.Round(1.7 * 1024 * 1024 * 1024)), true, "1.7GB"},

		// --- TB ---
		{"Human: 1TB exact", 1024 * 1024 * 1024 * 1024, true, "1TB"},
		{"Human: 3.2TB", int64(math.Round(3.2 * 1024 * 1024 * 1024 * 1024)), true, "3.2TB"},

		// --- PB ---
		{"Human: 1PB exact", 1024 * 1024 * 1024 * 1024 * 1024, true, "1PB"},
		{"Human: 5.5PB", int64(5.5 * 1024 * 1024 * 1024 * 1024 * 1024), true, "5.5PB"},

		// --- EB ---
		{"Human: 1EB exact", 1024 * 1024 * 1024 * 1024 * 1024 * 1024, true, "1EB"},
		{"Human: 2.1EB", int64(math.Round(2.1 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024)), true, "2.1EB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatSize(tt.size, tt.human)
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}
