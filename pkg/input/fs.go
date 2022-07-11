package input

import (
	"io/fs"
	"os"
	"path"
	"strings"

	"k8s.io/client-go/util/homedir"
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
	if strings.HasPrefix(name, "~") {
		name = path.Join(homedir.HomeDir(), name[1:])
	}
	if !strings.HasPrefix(name, "/") {
		name = path.Join("./", d.wd, name)
	}
	name = path.Clean("./" + name)

	return d.root.Open(name) //nolint:wrapcheck
}

func currentDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
