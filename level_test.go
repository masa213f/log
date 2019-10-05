package log

import (
	"testing"
)

func TestStringToLevel(t *testing.T) {
	testcase := []struct {
		input    []string
		expected Level
	}{
		{
			input:    []string{"d", "D", "debug", "DEBUG", "Debug", "deBuG"},
			expected: LevelDebug,
		},
		{
			input:    []string{"i", "I", "info", "INFO", "Info", "inFO"},
			expected: LevelInfo,
		},
		{
			input:    []string{"w", "W", "warn", "WARN", "warning", "Warning"},
			expected: LevelWarning,
		},
		{
			input:    []string{"e", "E", "err", "ERR", "error", "ERROR"},
			expected: LevelError,
		},
		{
			input:    []string{"", "def", "default", "hoge", "piyo", "fuga"},
			expected: LevelDefault,
		},
	}
	for _, tc := range testcase {
		for _, input := range tc.input {
			t.Run(input, func(t *testing.T) {
				lv := StringToLevel(input)
				if lv != tc.expected {
					t.Errorf("expected=%d, actual=%d", tc.expected, lv)
				}
			})
		}
	}
}
