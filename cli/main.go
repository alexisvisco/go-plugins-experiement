package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"test-go-plugins/cli/addon"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "cli is a flexible command line interface with addon management",
}

func main() {

	addon.RegisterAddons(rootCmd)

	execute()
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "addons",
		Short: "List of addons enabled",
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, a := range addon.List {
				fmt.Printf("%s: %s\n", a.Name, a.Description)
			}
			return nil
		},
	})
}
