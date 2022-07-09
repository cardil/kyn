package cli

import (
	"context"
	"errors"
	"io"

	"github.com/cardil/kyn/pkg/ns"
)

var (
	ErrUnexpected            = errors.New("unexpected")
	ErrNoNamespaceToRename   = errors.New("no namespace to rename")
	ErrNamespaceInvalid      = errors.New("invalid namespace")
	ErrCanHaveOnlyOneDefault = errors.New("can have only one default namespace")
	ErrInvalidFilePath       = errors.New("invalid file path")
	ErrNoFileToRename        = errors.New("no file to rename")
)

type Rename struct {
	Namespaces []string
	Files      []string
}

func (r Rename) Do(ctx context.Context, out io.Writer, in io.Reader) error {
	if err := r.validate(); err != nil {
		return err
	}
	if rr, err := r.parse(out, in); err != nil {
		return err
	} else {
		return rr.Perform(ctx)
	}
}

func (r Rename) validate() error {
	if len(r.Namespaces) == 0 {
		return ErrNoNamespaceToRename
	}
	if len(r.Files) == 0 {
		return ErrNoFileToRename
	}
	return nil
}

func (r Rename) parse(out io.Writer, in io.Reader) (ns.Rename, error) {
	return ns.Rename{
		Output: out,
	}, ErrUnexpected
}
