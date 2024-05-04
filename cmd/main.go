package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/drornir/cue-example-api/gen"
)

func main() {
	if err := run(); err != 0 {
		os.Exit(err)
	}
}

func run() (errCode int) {
	cmdName := filepath.Base(os.Args[0])
	defer func() {
		if x := recover(); x != nil {
			_, _ = fmt.Fprintf(os.Stderr, cmdName+": panic: %v\n", x)
			errCode = 2
		}
	}()

	genCmd := gen.MakeRootCommand()

	rootCommand := &cobra.Command{
		Use:          cmdName,
		SilenceUsage: true,
	}
	rootCommand.AddCommand(
		genCmd,
	)
	rootCommand.SetErrPrefix(cmdName + " error:")

	ctx := context.Background()

	if err := rootCommand.ExecuteContext(ctx); err != nil {
		return 1
	}
	return 0
}
