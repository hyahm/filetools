package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"path"
)

var old string
var new string
var newr string

func init() {
	replaceCmd.Flags().StringVarP(&file, "file", "", "", "file path")
	replaceCmd.Flags().StringVarP(&old, "old", "", "", "file path")
	replaceCmd.Flags().StringVarP(&new, "new", "", "", "file path")
	replaceCmd.Flags().StringVarP(&newr, "newr", "", "", "file path")
	rootCmd.AddCommand(replaceCmd)
}

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "edit file",
	Long:  "edit file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-------")
		replace()

	},
}

func replace() {
	fmt.Println(1111)
	fmt.Println(path.Dir("."))
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("----")
		log.Fatal(err)
	}
	newf, _ := os.Create(file + "-1.txt")

	wd := bufio.NewWriter(newf)
	defer f.Close()
	defer newf.Close()
	rd := bufio.NewReader(f)
	count := 1
	for {
		var changelist []byte
		line, err := rd.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println(count)
				//读取完成， 退出循环
				fmt.Println("input over")
				break
			}
			log.Fatal(err)
		}

		changelist = append(changelist, []byte("file '")...)
		changelist = append(changelist, line[:len(line)-1]...)
		changelist = append(changelist, []byte("'\n")...)
		wd.Write(changelist)
		count++
	}
	wd.Flush()
}
