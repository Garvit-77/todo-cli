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

// COMPLETECmd represents the COMPLETE command
var COMPLETECmd = &cobra.Command{
	Use:   "COMPLETE",
	Short: "A brief description of your command",
	Long:  `-`,
	Run: func(cmd *cobra.Command, args []string) {
		taskid := args[0]
		toogle(taskid)
	},
}

func toogle(taskid string) {

	file, err := os.Open("tasks.csv")
	if err != nil {
		fmt.Println("Can't open your file: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Can't Read the Csv: ", err)
	}

	updatedtasks := [][]string{}
	found := false
	for _, task := range tasks {
		if task[0] == taskid {
			task[2] = "True"
			found = true
		}
		updatedtasks = append(updatedtasks, task)
	}

	if !found {
		fmt.Println("Could not find your task id:", err)
		return
	}

	file, err = os.Create("tasks.csv")
	if err != nil {
		fmt.Println("Can't Create your File: ", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range updatedtasks {
		err = writer.Write(task)
		if err != nil {
			fmt.Println("Not able to update: ", err)
		}
	}

	fmt.Println("Hey!, Succesfully Completed the task .\n", taskid)
}

func init() {
	rootCmd.AddCommand(COMPLETECmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// COMPLETECmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// COMPLETECmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
