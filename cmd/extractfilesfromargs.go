package cmd

import (
	"os"
	"path/filepath"

	"github.com/CentaurWarchief/subbs/fs"
)

func ExtractFilesFromArgs(args []string) (files []string) {
	for _, arg := range args {
		fi, err := os.Stat(arg)

		if err != nil {
			continue
		}

		abs, err := filepath.Abs(arg)

		if err != nil {
			continue
		}

		if fi.IsDir() {
			files = append(files, fs.ReadDir(abs)...)
			continue
		}

		files = append(files, abs)
	}

	return files
}
