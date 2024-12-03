/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var editDescription bool

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var options []string

		tasks, _ := GetTasks()

		for _, t := range tasks {
			options = append(options, t.Name)
		}
		taskSelected, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

		task := selectedTask(tasks, taskSelected)

		var promptMessage string

		if editDescription {
			promptMessage = task.Description
		} else {
			promptMessage = task.Name
		}

		updatedText, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue(promptMessage).Show()

		if editDescription {
			task.Description = updatedText
		} else {
			task.Name = updatedText
		}

		task.Write()
		pterm.Println()
		pterm.Info.Printfln("Task has been update to %s", pterm.Green(task.Name))
	},
}

func selectedTask(tasks []Task, taskName string) Task {
	for _, t := range tasks {
		if t.Name == taskName {
			return t
		}
	}
	return Task{} // this should never occur
}

func init() {
	editCmd.Flags().BoolVarP(&editDescription, "description", "d", false, "Edit Description")
	rootCmd.AddCommand(editCmd)
}
