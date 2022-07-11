package input

import (
	"io"
)

type Stream struct {
	Source string
	io.Reader
}

func (f Stream) Name() string {
	return f.Source
}

func (f Stream) Read() ([]io.ReadCloser, error) {
	return []io.ReadCloser{
		io.NopCloser(f.Reader),
	}, nil
}
