package main

import (
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/thavlik/mirbase/pkg/store/init_store"
)

var searchArgs struct {
	db string
}

var searchCmd = &cobra.Command{
	Use:  "search",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := strings.Join(args, " ")
		store, err := init_store.InitStore(searchArgs.db, false)
		if err != nil {
			return errors.Wrap(err, "failed to open database")
		}
		defer store.Close()
		results, err := store.SearchMiRNAs(cmd.Context(), query, 10, 1)
		if err != nil {
			return errors.Wrap(err, "failed to search miRNAs")
		}
		indented, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			return errors.Wrap(err, "failed to marshal results")
		}
		fmt.Println(string(indented))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&searchArgs.db, "db", "d", defaultDbPath, "Path to the SQLite database")
}
