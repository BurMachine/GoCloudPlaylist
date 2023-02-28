package timeConverting

import (
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
