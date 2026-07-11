package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"runtime/debug"

	"github.com/suer/dtfmt/internal/input"
	"github.com/suer/dtfmt/internal/output"
)

func Run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("dtfmt", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {
		_, _ = fmt.Fprintln(stderr, "usage: dtfmt <file-path|unix-timestamp|datetime-string>")
		fs.PrintDefaults()
	}
	var showVersion bool
	fs.BoolVar(&showVersion, "version", false, "print version")
	fs.BoolVar(&showVersion, "v", false, "print version")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}

	if showVersion {
		_, _ = fmt.Fprintln(stdout, "dtfmt "+version())
		return 0
	}

	positional := fs.Args()
	if len(positional) != 1 {
		fs.Usage()
		return 2
	}

	arg := positional[0]
	result, err := input.Detect(arg)
	if err != nil {
		_, _ = fmt.Fprintln(stderr, "dtfmt: "+err.Error())
		return 1
	}

	doc := output.Build(arg, result)

	enc := json.NewEncoder(stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(doc); err != nil {
		_, _ = fmt.Fprintln(stderr, "dtfmt: "+err.Error())
		return 1
	}
	return 0
}

func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" {
		return "(devel)"
	}
	return info.Main.Version
}
