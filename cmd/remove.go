package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var removeCmd = &cobra.Command{
	Use: "remove",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			todoName := args[0]
			for i := 0; i < len(todos.Todos); i++ {
				if todos.Todos[i].Name == todoName {
					todos.Todos = append(todos.Todos[0:i], todos.Todos[i+1:]...)
					break
				}
			}
			viper.Set("todos", todos.Todos)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
