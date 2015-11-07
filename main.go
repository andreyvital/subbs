package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

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
			fmt.Printf("Directory: %s\n", abs)
			continue
		}

		fmt.Printf("File: %s\n", abs)
	}
}
