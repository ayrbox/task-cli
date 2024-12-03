/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/nrednav/cuid2"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func parseArgs(args []string) (string, string) {
	if len(args) == 1 {
		return args[0], ""
	}
	return args[0], args[1]
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name, description := parseArgs(args)

		t := Task{
			ID:          cuid2.Generate(),
			Name:        name,
			Description: description,
		}

		if err := t.Write(); err != nil {
			fmt.Printf("error : %v", err)
		}

		pterm.Info.Println("New task added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
