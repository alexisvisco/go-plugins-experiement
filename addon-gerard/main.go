package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"test-go-plugins/cli/addon"
)

var Addon = addon.Addon{
	Name:        "gerard",
	Version:     "1.0",
	Author:      "alexisvisco",
	Repository:  "github.com/alexisvisco/zaap-gerard",
	Description: "This addon is for doing something",

	RegisterCommands: registerCommands,
}

func registerCommands(rootCmd *cobra.Command) error {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "gerard",
		Short:   "Gerard should say something",
		Example: "cli gerard",
		Version: "1.0",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Gerard say: I am the gerard addon and I will be good with you.")
			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	})

	return nil
}
