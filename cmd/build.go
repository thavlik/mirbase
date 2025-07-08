package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/thavlik/mirbase/pkg/build"
	"github.com/thavlik/mirbase/pkg/store"
	"github.com/thavlik/mirbase/pkg/store/sql_store"
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
		output, err := initStore(buildArgs.output, true)
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

func initStore(path string, create bool) (store.Store, error) {
	query := []string{
		"_sync=FULL",
		"_journal_mode=PERSIST",
		"_auto_vacuum=FULL",
		"_foreign_keys=yes",
		"_defer_foreign_keys=no",
		"_case_sensitive_like=no",
		"cache=shared",
	}
	if create {
		query = append(query, "mode=rwc")
	} else {
		query = append(
			query,
			//"immutable=yes",
			"mode=rw", // ro
			"_query_only=yes",
		)
		info, err := os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("failed to stat database file: %v", err)
		}
		sizeInMb := info.Size() / 1000 / 1000
		if sizeInMb < 1 {
			return nil, fmt.Errorf("database file '%s' is too small: %d KiB", path, info.Size()/1000)
		}
	}
	db, err := sql.Open(
		"sqlite3",
		fmt.Sprintf(
			"file:%s?%s",
			path,
			strings.Join(query, "&"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	output, err := sql_store.NewSqlStore(db, create)
	if err != nil {
		return nil, fmt.Errorf("failed to create SQL store: %v", err)
	}
	return output, nil
}
