package fs

import (
	"os"
	"path/filepath"
)

func ReadDir(dir string) []string {
	var files []string

	err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil
	}

	return files
}
