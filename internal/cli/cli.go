package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"github.com/suer/dtfmt/internal/input"
	"github.com/suer/dtfmt/internal/output"
)

func Run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("dtfmt", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {
		fmt.Fprintln(stderr, "usage: dtfmt <file-path|unix-timestamp|datetime-string>")
	}
	if err := fs.Parse(args); err != nil {
		return 2
	}

	positional := fs.Args()
	if len(positional) != 1 {
		fs.Usage()
		return 2
	}

	arg := positional[0]
	result, err := input.Detect(arg)
	if err != nil {
		fmt.Fprintln(stderr, "dtfmt: "+err.Error())
		return 1
	}

	doc := output.Build(arg, result)

	enc := json.NewEncoder(stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(doc); err != nil {
		fmt.Fprintln(stderr, "dtfmt: "+err.Error())
		return 1
	}
	return 0
}
