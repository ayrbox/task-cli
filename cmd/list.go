/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := GetTasks()

		listData := pterm.TableData{
			{"Name", "Description", "Completed"},
		}

		for _, t := range tasks {
			listData = append(listData, []string{
				t.Name,
				t.Description,
				strconv.FormatBool(t.Completed),
			})
		}

		pterm.DefaultTable.WithHasHeader().WithData(listData).Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
