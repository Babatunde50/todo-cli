package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		td, err := json.MarshalIndent(todos.Todos, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(td))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
