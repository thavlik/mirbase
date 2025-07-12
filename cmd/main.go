package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mirbase",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("please choose a subcommand")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
