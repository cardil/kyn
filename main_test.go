package main_test

import (
	"bytes"
	"testing"

	main "github.com/cardil/kyn"
	"github.com/cardil/kyn/internal/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/wavesoftware/go-commandline"
)

func TestTheMain(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	defer func() {
		cmd.Opts = []commandline.Option{}
	}()
	var gotErr error
	cmd.Opts = []commandline.Option{
		commandline.WithCommand(func(cmd *cobra.Command) {
			cmd.SetArgs([]string{"--help"})
			cmd.SetOut(buf)
			cmd.SetErr(buf)
		}),
		commandline.WithErrorHandler(func(err error, _ *cobra.Command) bool {
			gotErr = err
			return false
		}),
	}
	main.RunMain()

	assert.Contains(t, buf.String(), "Usage:")
	assert.NoError(t, gotErr)
}
