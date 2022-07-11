package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wavesoftware/go-commandline"
)

var Opts []commandline.Option

func (a *App) configureFlags(r *cobra.Command) {
	fl := r.PersistentFlags()
	fl.StringSliceVarP(&a.Rename.Namespaces, "namespace", "n",
		[]string{}, "Namespace(s) to change to. You could specify a "+
			"specific namespace to replace with from=to syntax. Can be specified"+
			" multiple times.")
	fl.StringSliceVarP(&a.Rename.Files, "file", "f",
		[]string{"-"}, "A YAML file or directory with YAMLs to rename namespace"+
			" in. Can be specified multiple times. Use `-` to read from stdin.")
}
