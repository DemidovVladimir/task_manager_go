package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/DemidovVladimir/cli/db"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Add task: %s \n", task)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
