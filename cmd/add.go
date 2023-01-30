package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			todoName := args[0]
			description, _ := cmd.Flags().GetString("description")
			deadline, _ := cmd.Flags().GetString("deadline")
			todo := Todo{
				Name:        todoName,
				Description: description,
				Deadline:    deadline,
			}

			todos.Todos = append(todos.Todos, todo)
			viper.Set("todos", todos.Todos)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("Enter just one todo ")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	addCmd.Flags().String("description", "", "Todo description")
	addCmd.Flags().String("deadline", tomorrow.String(), "Deadline for the todo")
}
