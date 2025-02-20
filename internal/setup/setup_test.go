package setup

import "testing"

func TestConstants(t *testing.T) {
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"DataPath", DataPath, "/tmp/data"},
		{"SqliteFilename", SqliteFilename, "/tmp/data/sqlite.db"},
		{"SqlDriver", SqlDriver, "sqlite"},
		{"DefaultEncKey", DefaultEncKey, "abcdefghijklmnopABCDEFGHIJKLMNOP1234567890!@#$"},
		{"BlockConcat", BlockConcat, "%s^%s^%s^%s^%s^%d^%s"},
		{"DefaultAlgo", DefaultAlgo, "SHA512"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, tc.got)
			}
		})
	}
}
