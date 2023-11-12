package main

import (
	"context"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/xingbase/dashlinux"
	"github.com/xingbase/dashlinux/http"
)

// Build flags
var (
	version = ""
	commit  = ""
)

func main() {
	srv := http.Server{
		BuildInfo: dashlinux.BuildInfo{
			Version: version,
			Commit:  commit,
		},
	}

	parser := flags.NewParser(&srv, flags.Default)
	parser.ShortDescription = `Dashlinux`
	parser.LongDescription = `Options for Dashlinux`

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	srv.Serve(context.Background())
}
