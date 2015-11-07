package util

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func SrtPath(file string) string {
	return path.Join(
		filepath.Dir(file),
		fmt.Sprintf(
			"%s.srt",
			strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)),
		),
	)
}
