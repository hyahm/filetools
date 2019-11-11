package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	editCmd.Flags().StringVarP(&file, "file", "", "", "file path")
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit file",
	Long:  "edit file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(file)

	},
}
