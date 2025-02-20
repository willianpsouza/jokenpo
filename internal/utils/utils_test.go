package utils

import (
	"os"
	"testing"
)

func TestCheckPathExists(t *testing.T) {
	testDir := "test_dir"

	_ = os.RemoveAll(testDir)

	checkPathExists(testDir)

	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Errorf("Expected directory %s to be created, but it does not exist", testDir)
	}

	_ = os.RemoveAll(testDir)
}

func TestCheck(t *testing.T) {
	testDirs := []string{"test_dir1", "test_dir2"}

	for _, dir := range testDirs {
		_ = os.RemoveAll(dir)
	}

	Check(testDirs)

	for _, dir := range testDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Expected directory %s to be created, but it does not exist", dir)
		}
	}

	for _, dir := range testDirs {
		_ = os.RemoveAll(dir)
	}
}
