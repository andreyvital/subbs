package opensubtitles

import (
	"bytes"
	"encoding/binary"
	"os"
)

const (
	ChunkSize = 65536
)

// http://trac.opensubtitles.org/projects/opensubtitles/wiki/HashSourceCodes
func Hash(file *os.File) (hash uint64, err error) {
	fi, err := file.Stat()

	if err != nil {
		return
	}

	buf := make([]byte, ChunkSize*2)

	_, err = file.ReadAt(buf[:ChunkSize], 0)

	if err != nil {
		return
	}

	_, err = file.ReadAt(buf[ChunkSize:], fi.Size()-ChunkSize)

	if err != nil {
		return
	}

	var n [(ChunkSize * 2) / 8]uint64

	err = binary.Read(bytes.NewReader(buf), binary.LittleEndian, &n)

	if err != nil {
		return
	}

	for _, num := range n {
		hash += num
	}

	return hash + uint64(fi.Size()), nil
}
