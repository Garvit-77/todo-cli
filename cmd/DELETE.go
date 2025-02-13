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

// DELETECmd represents the DELETE command
var DELETECmd = &cobra.Command{
	Use:   "DELETE",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskid := args[0]
		delete(taskid)
	},
}

func delete(taskid string) {

	file, err := os.Open("tasks.csv")
	if err != nil {
		fmt.Println("Can't open the file: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Can' read your file:", err)
		return
	}

	updated := [][]string{}
	found := false

	for _, task := range tasks {
		if task[0] == taskid {
			continue
		}
		updated = append(updated, task)
	}

	if !found {
		fmt.Println("Can't find this task:", taskid)
	}

	file, err = os.Create("tasks.csv")
	if err != nil {
		fmt.Println("Could not Create the duplicate file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range updated {
		err = writer.Write(task)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}

	fmt.Println("Deletion Accomplished. \n", taskid)
}

func init() {
	rootCmd.AddCommand(DELETECmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DELETECmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DELETECmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
