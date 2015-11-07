package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/CentaurWarchief/subbs/fs"
	"github.com/CentaurWarchief/subbs/opensubtitles"
	"github.com/CentaurWarchief/subbs/util"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	var files []string

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

			files = append(
				files,
				fs.ReadDir(abs)...,
			)
		} else {
			fmt.Printf("File: %s\n", abs)

			files = append(
				files,
				abs,
			)
		}
	}

	client := opensubtitles.NewClient()

	if err := client.Authenticate("", ""); err != nil {
		return
	}

	var wg sync.WaitGroup

	wg.Add(len(files))

	for _, path := range files {
		go func(wg *sync.WaitGroup, path string) {
			defer wg.Done()

			file, err := os.Open(path)

			if err != nil {
				return
			}

			hash, err := opensubtitles.Hash(file)

			if err != nil {
				return
			}

			fmt.Printf("%16x %s\n", hash, filepath.Base(path))

			subtitles, err := client.Search(
				fmt.Sprintf("%x", hash),
				// TODO: add a flag that allows user to specify languages
				[]string{"eng", "pob"},
			)

			if err != nil {
				return
			}

			if len(subtitles) == 0 {
				// feedback
				return
			}

			subtitle, err := client.Download(*subtitles.First())

			if err != nil {
				return
			}

			reader, err := subtitle.Reader()

			if err != nil {
				return
			}

			defer reader.Close()

			writer, err := os.Create(util.SrtPath(path))

			if err != nil {
				return
			}

			defer writer.Close()

			_, err = io.Copy(writer, reader)
			// feedback
		}(&wg, path)
	}

	wg.Wait()
}
