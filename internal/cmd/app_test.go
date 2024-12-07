package cmd_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/cardil/kyn/internal/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wavesoftware/go-commandline"
)

func TestAppCommandHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	err := commandline.New(&cmd.App{}).Execute(commandline.WithCommand(func(cmd *cobra.Command) {
		cmd.SetArgs([]string{"--help"})
		cmd.SetOut(buf)
		cmd.SetErr(buf)
	}))
	assert.Contains(t, buf.String(), "Usage:")
	require.NoError(t, err)
}

func TestAppCommandRun(t *testing.T) {
	app := cmd.App{}
	c := app.Command()
	assert.Equal(t, "kyn", c.Use)

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs([]string{"--namespace", "acme", "--file", "./"})
	// Pre-cancelled context
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()
	c.SetContext(ctx)
	err := c.Execute()
	assert.ErrorIs(t, err, context.Canceled)
}
