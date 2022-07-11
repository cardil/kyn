package input

import (
	"io/fs"
	"os"
	"path"
)

func CurrentWorkingDirFS() fs.FS {
	return dirFS{
		wd:   currentDir(),
		root: os.DirFS("/"),
	}
}

type dirFS struct {
	wd   string
	root fs.FS
}

func (d dirFS) Open(name string) (fs.File, error) {
	p := path.Join("./", d.wd, name)
	p = path.Clean(p)
	return d.root.Open(p) //nolint:wrapcheck
}

func currentDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
