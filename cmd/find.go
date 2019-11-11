package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var file string

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().StringVarP(&file, "file", "", "", "filename")
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "find file or content",
	Long:  "find file or content",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file 8888", file)
	},
}

func F() {

}
