package timeConverting

import (
	"GoCloudPlaylist/internal/playlist"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertFromSecondsToString(t *testing.T) {
	testTable := []struct {
		seconds  int
		expected string
	}{
		{
			seconds:  12,
			expected: "00:00:12",
		},
		{
			seconds:  0,
			expected: "00:00:00",
		},
		{
			seconds:  -1,
			expected: "",
		},
		{
			seconds:  99999,
			expected: "27:46:39",
		},
	}

	for _, testCase := range testTable {
		result := ConvertFromSecondsToString(testCase.seconds)
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("incorrect result, expected %s, got %s",
			testCase.expected, result))
	}
}

func TestParseTimeToSeconds(t *testing.T) {
	testTable := []struct {
		time     string
		expected int
	}{
		{
			expected: 12,
			time:     "00:00:12",
		},
		{
			expected: 0,
			time:     "00:00:00",
		},
		{
			expected: 81999,
			time:     "22:46:39",
		},
	}

	for _, testCase := range testTable {
		result, err := ParseTimeToSeconds(testCase.time)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("incorrect result, expected %d, got %d",
			testCase.expected, result))
	}
	result, err := ParseTimeToSeconds("")
	if err == nil {
		t.Error(err)
	}
	assert.Equal(t, 0, result, fmt.Sprintf("incorrect result, expected %d, got %d",
		0, result))
}

func TestConvertFromSongProcToString(t *testing.T) {
	testTable := []struct {
		songProc playlist.SongProcessing
		expected string
	}{
		{
			playlist.SongProcessing{
				Name:        "string",
				Duration:    10,
				CurrentTime: 9,
				Playing:     false,
				Exist:       true,
			},
			"00:00:09 of 00:00:10",
		},
		{
			playlist.SongProcessing{
				Name:        "string",
				Duration:    70,
				CurrentTime: 60,
				Playing:     false,
				Exist:       true,
			},
			"00:01:00 of 00:01:10",
		},
		{
			playlist.SongProcessing{
				Name:        "string",
				Duration:    0,
				CurrentTime: 0,
				Playing:     false,
				Exist:       true,
			},
			"00:00:00 of 00:00:00",
		},
		{
			playlist.SongProcessing{
				Name:        "string",
				Duration:    -1,
				CurrentTime: -1,
				Playing:     false,
				Exist:       true,
			},
			" of ",
		},
	}

	for _, testCase := range testTable {
		result := ConvertFromSongProcToString(testCase.songProc)
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("incorrect result, expected %s, got %s",
			testCase.expected, result))
	}
}
