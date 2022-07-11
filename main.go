package main

import (
	"github.com/cardil/kyn/internal/cmd"
	"github.com/wavesoftware/go-commandline"
)

func main() {
	app := new(cmd.App)
	commandline.New(app).ExecuteOrDie(cmd.Opts...)
}
