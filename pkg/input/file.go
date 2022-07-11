package input

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
)

var ErrCouldntReadFile = errors.New("couldn't read file")

type File struct {
	fs.FS
	Path string
}

func (f File) Name() string {
	return f.Path
}

func (f File) Read() ([]io.ReadCloser, error) {
	file, err := f.FS.Open(f.Path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %v", ErrCouldntReadFile, f.Path, err)
	}
	return []io.ReadCloser{file}, nil
}
