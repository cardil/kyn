package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"

	"github.com/cardil/kyn/pkg/ns"
)

var (
	ErrUnexpected          = errors.New("unexpected")
	ErrCanHaveOnlyOneStdin = errors.New("can have only one stdin")
	ErrInvalidFilePath     = errors.New("invalid file path")
	StdinRepresentation    = "-"
)

type Rename struct {
	Namespaces []string
	Files      []string
	fs.FS
}

func (r Rename) Do(ctx context.Context, out io.Writer, in io.Reader) error {
	var (
		rename ns.Rename
		err    error
	)

	if rename, err = r.parse(out, in); err != nil {
		return err
	}
	if err = rename.Perform(ctx); err != nil {
		err = fmt.Errorf("%w: %v", ErrUnexpected, err)
	}
	return err
}
