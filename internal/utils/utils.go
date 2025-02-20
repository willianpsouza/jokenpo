package utils

import "os"

func checkPathExists(path string) {
	_, err := os.Stat(path)
	if err != nil {
		_ = os.MkdirAll(path, os.ModePerm)
	}
}

func Check(paths []string) {
	for _, path := range paths {
		checkPathExists(path)
	}
}
