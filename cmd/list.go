package cmd

import (
	"fmt"
	"os"

	"github.com/DemidovVladimir/cli/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List taksks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong...")
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You don't have any tasks...")
			return
		}

		fmt.Println("Here are the list of the tasks:")

		for i, task := range tasks {
			fmt.Printf("%d. %s \n", i+1, task.Value)
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
