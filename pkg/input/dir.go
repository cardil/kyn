package input

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

var ErrCouldntReadDir = errors.New("couldn't read dir")

const nonEmptyDirectoryCapacity = 10

type Dir struct {
	File
}

func (f Dir) Read() ([]io.ReadCloser, error) {
	files := make([]io.ReadCloser, 0, nonEmptyDirectoryCapacity)
	err := fs.WalkDir(f.FS, f.Path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".yaml") && !strings.HasSuffix(path, ".yml") {
			return nil
		}
		file, err := f.Open(path)
		if err != nil {
			return fmt.Errorf("%w: %s: %w", ErrCouldntReadFile, path, err)
		}

		files = append(files, file)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrCouldntReadDir, f.Path, err)
	}
	return files, nil
}
