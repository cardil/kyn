package ns

import (
	"context"
	"io"
)

const AnyNamespace = ""

type Rename struct {
	Namespaces []NamespaceRename
	Inputs     []Input
	Output     io.Writer
}

type NamespaceRename struct {
	From string
	To   string
}

type Input interface {
	Name() string
	Read() ([]io.ReadCloser, error)
}

func (r Rename) Perform(ctx context.Context) error {
	if err := r.Validate(); err != nil {
		return err
	}
	m, err := r.readManifest(ctx)
	if err != nil {
		return err
	}
	m, err = r.injectNamespace(m)
	if err != nil {
		return err
	}

	return r.output(ctx, m.Resources())
}
