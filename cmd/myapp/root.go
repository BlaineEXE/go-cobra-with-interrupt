package myapp

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "myapp",
	Short:         "My example app",
	RunE:          runCmd,
	SilenceUsage:  true,
	SilenceErrors: true, // already taken care of in ExecuteContext()
}

func ExecuteContext(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runCmd(cmd *cobra.Command, args []string) error {
	return thisTakesAContext(cmd.Context())
}

func thisTakesAContext(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf(" interrupted: %w", ctx.Err())
		case <-ticker.C:
			fmt.Println("tick", i)
			i++
		}
	}
}
