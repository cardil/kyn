package ns

import (
	"context"
	"io"
	"io/fs"
)

type Rename struct {
	Namespaces []NamespaceRename
	Files      []InputFile
	Output     io.Writer
}

func (r Rename) Perform(ctx context.Context) error {
	return nil
}

type NamespaceRename struct {
	From string
	To   string
}

type InputFile struct {
	fs.FS
	Path string
}
