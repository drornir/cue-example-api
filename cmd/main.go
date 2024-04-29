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
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	genCmd := gen.MakeRootCommand()

	cmdName := filepath.Base(os.Args[0])
	rootCommand := &cobra.Command{
		Use: cmdName,
	}
	rootCommand.AddCommand(
		genCmd,
	)

	ctx := context.Background()

	if err := rootCommand.ExecuteContext(ctx); err != nil {
		return fmt.Errorf("%s: %w", rootCommand.Name(), err)
	}
	return nil
}
