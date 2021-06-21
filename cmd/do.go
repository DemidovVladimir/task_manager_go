package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DemidovVladimir/cli/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Do task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Wrong agrument: %s \n", id)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong... ", err)
			os.Exit(1)
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("Invalid print number: %d", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Error when removing: \"%d\" - err: %s \n", id, err)
			} else {
				fmt.Printf("Removed: \"%d\" \n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
