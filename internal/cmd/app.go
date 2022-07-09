package cmd

import (
	"errors"

	"github.com/cardil/kyn/pkg/cli"
	"github.com/spf13/cobra"
)

type App struct {
	cli.Rename
}

func (a *App) Command() *cobra.Command {
	r := &cobra.Command{
		Use:   "kyn",
		Short: "Kubernetes YAML Namespace changer",
		Args:  cobra.NoArgs,
		Example: `
  kyn -n foo=bar -n acme -f ./yamls/ | kubectl apply -f -`,
		Long: "Kubernetes YAML Namespace changer - Change the namespace of " +
			"Kubernetes YAMLs and output modified files to stdout.",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := a.Do(cmd.Context(), cmd.OutOrStdout(), cmd.InOrStdin())
			cmd.SilenceUsage = errors.Is(err, cli.ErrUnexpected)
			return err
		},
	}
	a.configureFlags(r)
	return r
}

func (a App) run(cmd *cobra.Command) error {
	err := a.Do(cmd.Context(), cmd.OutOrStdout(), cmd.InOrStdin())
	cmd.SilenceUsage = errors.Is(err, cli.ErrUnexpected)
	return err
}
