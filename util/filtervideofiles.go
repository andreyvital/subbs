package util

import "path/filepath"

var SupportedVideoFileExtensions = [...]string{
	".avi",
	".mkv",
	".mp4",
}

func FilterVideoFiles(files []string) (f []string) {
	for _, file := range files {
		ext := filepath.Ext(file)

		for _, supported := range SupportedVideoFileExtensions {
			if ext == supported {
				f = append(f, file)
				break
			}
		}
	}

	return f
}
