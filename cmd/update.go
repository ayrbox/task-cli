/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var options []string

		tasks, _ := GetTasks()

		for _, t := range tasks {
			options = append(options, t.Name)
		}

		taskSelected, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

		pterm.Printfln("Do you want to mark the task as completed?")
		completed, _ := pterm.DefaultInteractiveConfirm.Show()

		for _, t := range tasks {
			if t.Name == taskSelected {
				t.Completed = completed
				t.Write()
			}
		}

		pterm.Info.Printfln("Task %s has been marked as %s", pterm.Yellow(taskSelected), boolToCompleted(completed))
	},
}

func boolToCompleted(b bool) string {
	if b {
		return pterm.Green("Completed")
	}
	return pterm.Red("Not Completed")
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
