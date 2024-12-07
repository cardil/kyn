package cli

import (
	"fmt"
	"io"
	"io/fs"
	"strings"

	"github.com/cardil/kyn/pkg/input"
	"github.com/cardil/kyn/pkg/ns"
)

const partsOfEquals = 2

func (r Rename) parse(out io.Writer, in io.Reader) (ns.Rename, error) {
	var err error
	rn := ns.Rename{
		Output:     out,
		Namespaces: r.parseNamespaces(),
	}
	if rn.Inputs, err = r.parseFiles(in); err != nil {
		return ns.Rename{}, err
	}
	err = rn.Validate()
	if err != nil {
		return ns.Rename{}, err //nolint:wrapcheck
	}
	return rn, nil
}

func (r Rename) parseNamespaces() []ns.NamespaceRename {
	names := make([]ns.NamespaceRename, 0, len(r.Namespaces))
	for _, namespacePair := range r.Namespaces {
		parts := strings.SplitN(namespacePair, "=", partsOfEquals)
		to := parts[0]
		from := ns.AnyNamespace
		if len(parts) > 1 {
			to, from = parts[0], parts[1]
		}
		names = append(names, ns.NamespaceRename{
			From: from,
			To:   to,
		})
	}
	return names
}

func (r Rename) parseFiles(in io.Reader) ([]ns.Input, error) {
	stdinPresent := false
	files := make([]ns.Input, 0, len(r.Files))
	for _, file := range r.Files {
		inpt, err := r.toInput(file, in)
		if err != nil {
			return nil, err
		}
		if _, ok := inpt.(input.Stream); ok {
			if stdinPresent {
				return nil, ErrCanHaveOnlyOneStdin
			}
			stdinPresent = true
		}
		files = append(files, inpt)
	}
	return files, nil
}

func (r Rename) toInput(file string, in io.Reader) (ns.Input, error) {
	if file == StdinRepresentation {
		return input.Stream{
			Reader: in,
			Source: "stdin",
		}, nil
	}
	fi, err := fs.Stat(r.FS, file)
	if err != nil {
		return nil, fmt.Errorf("%w: `%s`: %w", ErrInvalidFilePath, file, err)
	}
	if fi.IsDir() {
		return input.Dir{File: input.File{FS: r.FS, Path: file}}, nil
	}
	return input.File{FS: r.FS, Path: file}, nil
}
