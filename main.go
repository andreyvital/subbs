package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/CentaurWarchief/subbs/cmd"
	"github.com/CentaurWarchief/subbs/config"
	"github.com/CentaurWarchief/subbs/opensubtitles"
	"github.com/CentaurWarchief/subbs/util"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	files := util.FilterVideoFiles(cmd.ExtractFilesFromArgs(args))

	var config config.Config

	wd, _ := os.Getwd()

	if _, err := toml.DecodeFile(
		filepath.Join(wd, "config.toml"),
		&config,
	); err != nil {
		return
	}

	client := opensubtitles.NewClient()

	if err := client.Authenticate(
		config.OS.User,
		config.OS.Password,
	); err != nil {
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
				config.OS.Languages,
			)

			if err != nil || len(subtitles) == 0 {
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

			srt := util.SrtPath(path)

			writer, err := os.Create(srt)

			if err != nil {
				return
			}

			defer writer.Close()

			_, err = io.Copy(writer, reader)

			if err != nil {
				return
			}

			fmt.Println(filepath.Base(srt))
		}(&wg, path)
	}

	wg.Wait()
}
