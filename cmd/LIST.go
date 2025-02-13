/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// LISTCmd represents the LIST command
var LISTCmd = &cobra.Command{
	Use:   "LIST",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func list()  {
	file , err := os.Open("tasks.csv")
	if err != nil {
		fmt.Println("Could not open your file")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Unable to read file")
		return
	}

	if len(tasks) < 1 {
		fmt.Println("There are no tasks to read")
	} else {
		fmt.Println("TODO LIST")
		for _, task := range tasks {
			fmt.Printf("%s: %s [%s]\n", task[0], task[1], task[2])
		}
	}
}

func init() {
	rootCmd.AddCommand(LISTCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// LISTCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// LISTCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
