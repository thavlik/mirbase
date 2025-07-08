package main

import (
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/thavlik/mirbase/pkg/build"
	"github.com/thavlik/mirbase/pkg/store/init_store"
)

const defaultDbPath = "/mirbase.sqlite"

var buildArgs struct {
	output string
}

var buildCmd = &cobra.Command{
	Use:  "build",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()
		fmt.Println("Creating output database:", buildArgs.output)
		output, err := init_store.InitStore(buildArgs.output, true)
		if err != nil {
			return fmt.Errorf("failed to initialize store: %v", err)
		}
		defer func() {
			if err := output.Close(); err != nil {
				fmt.Printf("Warning: failed to close output db: %v\n", err)
				panic(err)
			}
		}()
		fmt.Println("All files opened. Building database...")
		if err := build.Build(cmd.Context(), output); err != nil {
			fmt.Println(err)
			return fmt.Errorf("failed to build database: %v", err)
		}
		fmt.Printf("Completed in %.2f seconds\n", time.Since(start).Seconds())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringVarP(&buildArgs.output, "output", "o", defaultDbPath, "Output SQLite database path")
}
